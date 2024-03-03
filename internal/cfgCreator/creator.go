package cfgCreator

import (
	"gopkg.in/yaml.v3"
	"ingressPlugin/internal/domain"
	"log"
	"os"
)

type CfgCreator struct {
	logger *log.Logger
}

func New(logger *log.Logger) *CfgCreator {
	return &CfgCreator{logger: logger}
}
func (c *CfgCreator) ChangeConfig(cfg *domain.Configuration) error {
	marshaledCfg, err := yaml.Marshal(cfg)
	if err != nil {
		c.logger.Printf("%s", err.Error())
		return err
	}

	err = os.WriteFile(cfg.Providers.File.Filename, marshaledCfg, 0666)
	if err != nil {
		c.logger.Printf("%s", err.Error())
		return err
	}

	return nil
}
