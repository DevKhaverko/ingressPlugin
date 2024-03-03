package nomad

import (
	"github.com/hashicorp/nomad/api"
	"ingressPlugin/internal/domain"
	"log"
)

type NomadIngressGetter struct {
	nomadClient *api.Client
	logger      *log.Logger
}

func New(
	nomadClient *api.Client,
	logger *log.Logger,
) *NomadIngressGetter {
	return &NomadIngressGetter{
		nomadClient: nomadClient,
		logger:      logger,
	}
}

func (n *NomadIngressGetter) GetIngressDefinition(i domain.IngressResource) (map[string]string, string, error) {
	info, _, err := n.nomadClient.Allocations().Info(i.GetDefinition(), nil)
	if err != nil {
		n.logger.Printf("%s", err.Error())
		return nil, "", err
	}

	return info.Job.Meta, info.NetworkStatus.Address, nil
}
