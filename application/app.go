package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/mmvergara/go-simple-api/routes"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	redisDb *redis.Client
}

func New() *App {
	app := &App{
		router: routes.LoadRoutes(),
		redisDb: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_DB_URL"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}
	
	fmt.Println("Connecting to redis")
	err := a.redisDb.Ping(ctx).Err()

	if err != nil {
		return fmt.Errorf("failed to connect to redis %w",err)
	}

	fmt.Println("Startin Server")
	err = server.ListenAndServe()

	if err != nil{
		return fmt.Errorf("failed to start server %w", err)
	}

	return nil
}