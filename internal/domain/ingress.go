package domain

type IngressResource interface {
	GetDefinition() string
}

type NomadIngressResource struct {
	AllocID string
}

func (n *NomadIngressResource) GetDefinition() string {
	return n.AllocID
}
