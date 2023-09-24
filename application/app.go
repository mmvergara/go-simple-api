package application

import (
	"context"
	"fmt"
	"net/http"
	"github.com/mmvergara/go-simple-api/routes"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: routes.LoadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil{
		return fmt.Errorf("failed to start server %w", err)
	}

	return nil
}