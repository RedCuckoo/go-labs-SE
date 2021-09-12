package handlers

import (
	"context"
	"github.com/deliveroo/jsonrpc-go"
	"github.com/redcuckoo/go-labs-SE/lab4/database"
	"github.com/redcuckoo/go-labs-SE/lab4/resources"
	"github.com/redcuckoo/go-labs-SE/lab4/service/helpers"
	"github.com/redcuckoo/go-labs-SE/lab4/service/requests"
	"log"
)

func UpdateCities(ctx context.Context, params resources.Params) (interface{}, error) {
	country, name,population,capital, err := requests.GetInsertCity(params)

	if err != nil {
		log.Println("Failed to get info from request")
		return nil, jsonrpc.Error("-1", "Failed to get info from request")
	}

	cityDB := helpers.CityDBFromCtx(ctx)

	city, err := cityDB.Insert(database.City{
		Country:    country,
		Name:       name,
		Population: population,
		Capital:    capital,
	})

	if err != nil {
		log.Printf("Failed to update, %v", err)
		return nil, jsonrpc.Error("-1","Failed to update info in database")
	}

	return city, nil
}
