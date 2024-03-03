package app

import (
	nomadApi "github.com/hashicorp/nomad/api"
	grpcapp "ingressPlugin/internal/app/grpc"
	"ingressPlugin/internal/cfgCreator"
	"ingressPlugin/internal/nomad"
	"ingressPlugin/internal/parser"
	"ingressPlugin/internal/service"
	"log"
)

type App struct {
	GRPC   *grpcapp.App
	logger *log.Logger
}

func New(logger *log.Logger) *App {
	client, err := nomadApi.NewClient(
		&nomadApi.Config{
			Address: "http://localhost:4646",
		})
	if err != nil {
		logger.Fatal(err)
		return nil
	}

	definitionParser := parser.New(logger)
	configCreator := cfgCreator.New(logger)
	nomadIngressGetter := nomad.New(client, logger)
	ingressService := service.New(
		logger,
		configCreator,
		nomadIngressGetter,
		definitionParser,
	)
	
	return &App{
		GRPC:   grpcapp.New(logger, ingressService),
		logger: logger,
	}
}
