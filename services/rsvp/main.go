package main

import (
	"github.com/pattyjogal/api/common/middleware"
	"github.com/pattyjogal/api/services/rsvp/config"
	"github.com/pattyjogal/api/services/rsvp/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/rsvp"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.RSVP_PORT, router))
}
