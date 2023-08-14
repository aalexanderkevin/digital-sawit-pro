package main

import (
	"digital-sawit-pro/config"
	"digital-sawit-pro/generated"
	"digital-sawit-pro/handler"
	"digital-sawit-pro/repository"
	"digital-sawit-pro/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.Instance()
	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(cfg.Service.Host + ":" + cfg.Service.Port))
}

func newServer() *handler.Server {
	db := storage.GetPostgresDb()

	userRepo := repository.NewUserRepository(db)

	opts := handler.NewServerOptions{
		Repository: userRepo,
	}
	return handler.NewServer(opts)
}
