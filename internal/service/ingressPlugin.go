package service

import (
	"context"
	"ingressPlugin/internal/domain"
	_ "ingressPlugin/proto/api"
	"log"
)

type ConfigCreator interface {
	ChangeConfig(*domain.Configuration) error
}

type IngressResourceGetter interface {
	GetIngressDefinition(domain.IngressResource) (map[string]string, string, error)
}

type Parser interface {
	CreateRouter(map[string]string) (*domain.Router, error)
	CreateService(string) (*domain.Service, error)
}

type IngressService struct {
	logger                     *log.Logger
	cfgCreator                 ConfigCreator
	nomadIngressResourceGetter IngressResourceGetter
	parser                     Parser
	cfg                        *domain.Configuration
}

func New(
	logger *log.Logger,
	creator ConfigCreator,
	getter IngressResourceGetter,
	definitionParser Parser,
) *IngressService {
	return &IngressService{
		logger:                     logger,
		cfgCreator:                 creator,
		nomadIngressResourceGetter: getter,
		parser:                     definitionParser,
		cfg:                        domain.GetDefaultConfiguration(),
	}
}

func (i *IngressService) CreateOrChangeRoute(ctx context.Context, allocID string) (bool, error) {

	definition, address, err := i.nomadIngressResourceGetter.GetIngressDefinition(&domain.NomadIngressResource{AllocID: allocID})
	if err != nil {
		i.logger.Printf("%s", err.Error())
		return false, err
	}

	err = i.updateCfg(allocID, definition, address, i.cfg)
	if err != nil {
		i.logger.Printf("%s", err.Error())
		return false, err
	}

	err = i.cfgCreator.ChangeConfig(i.cfg)
	if err != nil {
		i.logger.Printf("%s", err.Error())
		return false, err
	}

	return true, nil
}

func (i *IngressService) updateCfg(allocID string, definition map[string]string, address string, cfg *domain.Configuration) error {

	router, err := i.parser.CreateRouter(definition)
	if err != nil {
		i.logger.Printf("%s", err.Error())
		return err
	}

	service, err := i.parser.CreateService(address)
	if err != nil {
		i.logger.Printf("%s", err.Error())
		return err
	}

	cfg.HTTP.Routers[allocID+"router"] = router
	cfg.HTTP.Services[allocID+"service"] = service

	return nil
}
