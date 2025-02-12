package main

import (
	"context"
	"fmt"
	"github.com/k1v4/url_shortener/internal/config"
	"github.com/k1v4/url_shortener/internal/repository/postgres_repo"
	"github.com/k1v4/url_shortener/internal/service"
	"github.com/k1v4/url_shortener/internal/transport/grpc"
	"github.com/k1v4/url_shortener/pkg/database/postgres"
	"github.com/k1v4/url_shortener/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	authLogger := logger.New(logger.ServiceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, authLogger)

	cfg := config.New()
	if cfg == nil {
		panic("load config fail")
	}

	authLogger.Info(ctx, "read config successfully")
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBConfig.UserName,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.DbName,
	)

	pg, err := postgres.New(url, postgres.MaxPoolSize(cfg.DBConfig.PoolMax))
	if err != nil {
		panic(err)
	}

	linksRepository := postgres_repo.NewLinksRepository(pg)

	linksServ := service.NewLinksService(linksRepository)

	grpcServer, err := grpc.NewServer(ctx, cfg.GRPCServerPort, cfg.RestServerPort, linksServ)
	if err != nil {
		authLogger.Error(ctx, err.Error())
		return
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	// запуск сервера
	go func() {
		if err = grpcServer.Start(ctx); err != nil {
			authLogger.Error(ctx, err.Error())
		}
	}()

	<-graceCh

	err = grpcServer.Stop(ctx)
	if err != nil {
		authLogger.Error(ctx, err.Error())
	}
	authLogger.Info(ctx, "Server stopped")
}
