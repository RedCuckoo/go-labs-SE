package service

import (
	"encoding/json"
	"github.com/google/jsonapi"
	"github.com/redcuckoo/go-labs-SE/lab3/database"
	"net/http"
)

var cityDB database.CityDB

func getCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	cities, err := cityDB.Select()
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(cities)

	if err != nil {
		return
	}
}

func insertCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var city database.City

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		return
	}

	insert, err := cityDB.Insert(city)
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(insert)

	if err != nil {
		return
	}
}

func updateCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var city database.City

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		return
	}

	update, err := cityDB.Update(city)
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(update)

	if err != nil {
		return
	}
}

func deleteCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var city database.City

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		return
	}

	err = cityDB.Delete(city)
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode("Successful")

	if err != nil {
		return
	}
}
