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
		http.Error(w, "Failed to connect to database", 500)
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
		http.Error(w, "Failed to decode input", 500)
		return
	}

	insert, err := cityDB.Insert(city)
	if err != nil {
		http.Error(w, "Failed to connect to database", 500)
		return
	}

	err = json.NewEncoder(w).Encode(insert)

	if err != nil {
		http.Error(w, "Failed to encode body", 500)
		return
	}
}

func updateCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var city database.City

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		http.Error(w, "Failed to decode body", 500)
		return
	}

	update, err := cityDB.Update(city)
	if err != nil {
		http.Error(w, "Failed to connect to database", 500)
		return
	}

	err = json.NewEncoder(w).Encode(update)

	if err != nil {
		http.Error(w, "Failed to encode body", 500)
		return
	}
}

func deleteCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var city database.City

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		http.Error(w, "Failed to decode body", 500)
		return
	}

	err = cityDB.Delete(city)
	if err != nil {
		http.Error(w, "Failed to connect to database", 500)
		return
	}

	err = json.NewEncoder(w).Encode("Successful")

	if err != nil {
		http.Error(w, "Failed to encode body", 500)
		return
	}
}
