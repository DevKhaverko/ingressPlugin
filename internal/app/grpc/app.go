package grpc

import (
	"google.golang.org/grpc"
	"ingressPlugin/internal/grpc/ingress"
	"log"
	"net"
)

type App struct {
	gRPC   *grpc.Server
	logger *log.Logger
}

func New(logger *log.Logger, ingressService ingress.IngressPluginService) *App {
	gRPCServer := grpc.NewServer()
	ingress.RegisterServerAPI(gRPCServer, ingressService)

	return &App{
		gRPC:   gRPCServer,
		logger: logger,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", ":777")
	if err != nil {
		a.logger.Printf("%s", err.Error())
		return err
	}
	a.logger.Printf("Listen on port")
	err = a.gRPC.Serve(l)
	if err != nil {
		a.logger.Printf("%s", err.Error())
		return err
	}
	a.logger.Printf("GRPC server started")
	return nil
}

func (a *App) Stop() {
	a.gRPC.GracefulStop()
}
