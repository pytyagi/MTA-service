package main

import (
	"log"
	api "mta-hosting-optimizer/api"
	"mta-hosting-optimizer/application"
	"mta-hosting-optimizer/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load the application configs
	_ = application.LoadConfiguration()
	// Initiated a mux router
	router := mux.NewRouter()
	ipDataSvc := service.NewIpDataService(&http.Client{})
	handler := api.NewListingHandler(ipDataSvc)
	// routes for getting ineffcient servers info
	router.HandleFunc("/hosts", handler.ListInefficientServersHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
