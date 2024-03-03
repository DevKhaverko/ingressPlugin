package parser

import (
	"ingressPlugin/internal/domain"
	"log"
)

type TraefikParser struct {
	logger *log.Logger
}

func New(logger *log.Logger) *TraefikParser {
	return &TraefikParser{logger: logger}
}

func (p *TraefikParser) CreateRouter(definition map[string]string) (*domain.Router, error) {
	r := &domain.Router{
		Service: definition["sevice_name"],
		Rule:    definition["rule"],
	}
	return r, nil
}

func (p *TraefikParser) CreateService(address string) (*domain.Service, error) {
	r := &domain.Service{
		LoadBalancer: &domain.ServersLoadBalancer{
			Servers: []domain.Server{
				{
					URL: address,
				},
			},
		},
	}
	return r, nil
}
