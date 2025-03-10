package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zandala88/gorder/internal/common/config"
	"github.com/zandala88/gorder/internal/common/genproto/stockpb"
	"github.com/zandala88/gorder/internal/common/server"
	"github.com/zandala88/gorder/stock/ports"
	"github.com/zandala88/gorder/stock/service"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatalf("could not read config: %v", err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serviceType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)
	switch serviceType {
	case "grpc":
		server.RunRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stdockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":

	default:
		panic("Unknown service type")
	}

}
