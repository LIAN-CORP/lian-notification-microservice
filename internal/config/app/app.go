package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lian-corp/notification-microservice/internal/config"
	"github.com/lian-corp/notification-microservice/internal/config/bootstrap"
)

func Run() {
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())

	go func ()  {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<- sigCh
		cancel()
	}()

	start(ctx, cfg)
}

func start(ctx context.Context, cfg *config.Config) {
	container := bootstrap.NewContainer(cfg)
	go container.NotificationConsumer.StartConsumer(ctx)
	<-ctx.Done()
}
