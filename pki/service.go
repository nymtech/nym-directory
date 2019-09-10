package pki

type service struct{}

// Service defines the REST service interface for PKI.
type Service interface {
	CreateNode() error
}

func newService(cfg *Config) *service {
	return &service{}
}

func (service *service) CreateNode() error {
	return nil
}
