package main

import (
	"context"
	"digital-sawit-pro/config"
	"digital-sawit-pro/container"
	"digital-sawit-pro/generated"
	"digital-sawit-pro/handler"
	"digital-sawit-pro/helper"
	"digital-sawit-pro/repository"
	"digital-sawit-pro/storage"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	cfg := config.Instance()
	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(cfg.Service.Host + ":" + cfg.Service.Port))
}

func newServer() *handler.Server {
	ctx := helper.ContextWithRequestId(context.Background(), ksuid.New().String())
	appProvider := &defaultAppProvider{}
	app, closeResourcesFn, err := appProvider.BuildContainer(ctx, buildOptions{
		Postgres: true,
	})
	if err != nil {
		panic(err)
	}
	if closeResourcesFn != nil {
		defer closeResourcesFn()
	}

	opts := handler.NewServerOptions{
		Repository: app.UserRepo(),
	}
	return handler.NewServer(opts)
}

type AppProvider interface {
	BuildContainer(ctx context.Context, options buildOptions) (*container.Container, func(), error)
}

type buildOptions struct {
	Postgres bool
}

type defaultAppProvider struct {
}

func (defaultAppProvider) BuildContainer(ctx context.Context, options buildOptions) (*container.Container, func(), error) {
	var db *gorm.DB
	cfg := config.Instance()

	appContainer := container.NewContainer()
	appContainer.SetConfig(cfg)

	if options.Postgres {
		db = storage.GetPostgresDb()
		appContainer.SetDb(db)

		userRepo := repository.NewUserRepository(db)
		appContainer.SetUserRepo(userRepo)
	}

	deferFn := func() {
		if db != nil {
			storage.CloseDB(db)
		}
	}

	return appContainer, deferFn, nil
}
