package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/k1v4/url_shortener/internal/config"
	"github.com/k1v4/url_shortener/internal/repository/in_memory"
	"github.com/k1v4/url_shortener/internal/repository/postgres_repo"
	"github.com/k1v4/url_shortener/internal/service"
	"github.com/k1v4/url_shortener/internal/transport/grpc"
	"github.com/k1v4/url_shortener/pkg/database/postgres"
	"github.com/k1v4/url_shortener/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

const (
	inMemory   = "in_memory"
	postgresDb = "postgres"
)

func main() {
	ctx := context.Background()

	// создаём логгер и добавляем его в контекст
	shortenerLogger := logger.New(logger.ServiceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, shortenerLogger)

	// инициализируем конфиг
	// при его отсутствии нет смысла дальше продолжать
	cfg := config.New()
	if cfg == nil {
		panic("load config fail")
	}

	shortenerLogger.Info(ctx, "read config successfully")

	// получаем флаг для определения хранилища
	dbFlag := flag.String("db", inMemory, "database connection flag")
	flag.Parse()

	// определяем хранилище
	var linksRepository service.ILinksRepository
	if *dbFlag == postgresDb {
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

		linksRepository = postgres_repo.NewLinksRepository(pg)

		shortenerLogger.Info(ctx, "using PostgreSQL")
	} else {
		linksRepository = in_memory.NewLinksRepository()

		shortenerLogger.Info(ctx, "using in-memory")
	}

	// создаём сервис
	linksServ := service.NewLinksService(linksRepository)

	// поднимает сервер
	grpcServer, err := grpc.NewServer(ctx, cfg.GRPCServerPort, cfg.RestServerPort, linksServ)
	if err != nil {
		shortenerLogger.Error(ctx, err.Error())
		return
	}

	// добавляем канал для отваливания сигналов graceful shutdown
	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	// запуск сервера
	go func() {
		if err = grpcServer.Start(ctx); err != nil {
			shortenerLogger.Error(ctx, err.Error())
		}
	}()

	// освобождаем канал при сигнале
	<-graceCh

	// останавливаем сервера и завершаем работу
	err = grpcServer.Stop(ctx)
	if err != nil {
		shortenerLogger.Error(ctx, err.Error())
	}
	shortenerLogger.Info(ctx, "Server stopped")
}
