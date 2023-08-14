package test

import (
	"digital-sawit-pro/generated"
	"digital-sawit-pro/handler"
	"net/http"

	"testing"

	"github.com/labstack/echo/v4"
)

func SetupHttpHandler(t *testing.T, opt handler.NewServerOptions) http.Handler {
	e := echo.New()
	var server generated.ServerInterface = handler.NewServer(opt)

	generated.RegisterHandlers(e, server)
	return e.Server.Handler
}
