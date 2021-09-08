package service

import (
	"fmt"
	"github.com/google/jsonapi"
	"github.com/redcuckoo/go-labs-SE/lab3/config"
	"github.com/redcuckoo/go-labs-SE/lab3/database/operations"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("text", jsonapi.MediaType)

	fmt.Fprint(w,"Welcome to the HomePage!\n" +
		"/countries - get all countries\n" +
		"/countries/insert - insert country to database\n" +
		"/countries/update - update country in database\n" +
		"/countries/delete - delete country from database\n" +
		"/cities - get all cities\n" +
		"/cities/insert - insert city to database\n" +
		"/cities/update - update city in database\n" +
		"/cities/delete - delete city from database\n")
}

func handleRequests() {
	configure, err := config.Configure()

	if err != nil {
		return
	}
	countryDB = operations.NewCountryDB(configure.DB)
	cityDB = operations.NewCityDB(configure.DB)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/countries", getCountries)
	http.HandleFunc("/countries/insert", insertCountry)
	http.HandleFunc("/countries/update", updateCountry)
	http.HandleFunc("/countries/delete", deleteCountry)
	http.HandleFunc("/cities", getCities)
	http.HandleFunc("/cities/insert", insertCity)
	http.HandleFunc("/cities/update", updateCity)
	http.HandleFunc("/cities/delete", deleteCity)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func Main() {
	handleRequests()
}
