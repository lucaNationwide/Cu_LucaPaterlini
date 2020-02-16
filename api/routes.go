package api

import (
	"Cu_LucaPaterlini/api/handlers"
	"Cu_LucaPaterlini/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
// ListenHTTP list the routes provided by the package and start the server
func ListenHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/exchange/{ticker:[A-Z]+}", handlers.HandlerExchange).Methods("GET")

	srv := &http.Server{
		Addr:         config.DefaultAddr,
		WriteTimeout: config.ServerWriteTimeout,
		ReadTimeout:  config.ServerReadTimeout,
		Handler:      router,
	}
	log.Fatal(srv.ListenAndServe())
}
