package handlers

import (
	"context"
	"github.com/deliveroo/jsonrpc-go"
	"github.com/redcuckoo/go-labs-SE/lab4/resources"
	"github.com/redcuckoo/go-labs-SE/lab4/service/helpers"
	"log"
)

func GetCities(ctx context.Context, params resources.Params) (interface{}, error) {
	cityDB := helpers.CityDBFromCtx(ctx)

	cities, err := cityDB.Select()

	if err != nil {
		log.Printf("Failed to get info, %v", err)
		return nil, jsonrpc.Error("-1","Failed to get info from database")
	}

	if len(cities) == 0{
		return nil, jsonrpc.Error("-1","No cities stored in database")
	}

	return cities, nil
}
