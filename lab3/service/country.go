package service

import (
	"encoding/json"
	"github.com/google/jsonapi"
	"github.com/redcuckoo/go-labs-SE/lab3/database"
	"net/http"
)

var countryDB database.CountryDB

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	countries, err := countryDB.Select()
	if err != nil {
		http.Error(w, "Failed to connect to database", 500)
		return
	}

	err = json.NewEncoder(w).Encode(countries)

	if err != nil {
		http.Error(w, "Failed to encode body", 500)
		return
	}
}

func insertCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var country database.Country

	err := json.NewDecoder(r.Body).Decode(&country)
	if err != nil {
		http.Error(w, "Failed to decode body", 500)
		return
	}

	insert, err := countryDB.Insert(country)
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

func updateCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var country database.Country

	err := json.NewDecoder(r.Body).Decode(&country)
	if err != nil {
		http.Error(w, "Failed to decode body", 500)
		return
	}

	update, err := countryDB.Update(country)
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

func deleteCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonapi.MediaType)

	var country database.Country

	err := json.NewDecoder(r.Body).Decode(&country)
	if err != nil {
		http.Error(w, "Failed to decode body", 500)
		return
	}

	err = countryDB.Delete(country)
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
