package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: getEnv("SERVER_PORT", ":3000"),
		Handler: a.router,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("%w",err)
	}

	return nil
}


func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}