package service

import (
	"github.com/deliveroo/jsonrpc-go"
	"github.com/redcuckoo/go-labs-SE/lab4/config"
	"github.com/redcuckoo/go-labs-SE/lab4/database/operations"
	"github.com/redcuckoo/go-labs-SE/lab4/service/handlers"
	"github.com/redcuckoo/go-labs-SE/lab4/service/helpers"
	"net/http"
)

func NewRouter(cfg config.Config) http.Handler {
	server := jsonrpc.New()

	server.Use(
		CtxMiddleware(
			helpers.CtxCountryDB(operations.NewCountryDB(cfg.DB)),
			helpers.CtxCityDB(operations.NewCityDB(cfg.DB)),
		),
	)

	server.Register(jsonrpc.Methods{
		"getcities": handlers.GetCities,
		"insertcity": handlers.InsertCities,
		"updatecity": handlers.UpdateCities,
	})

	return server
}
