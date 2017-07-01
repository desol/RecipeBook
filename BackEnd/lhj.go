package main

import (
	"log"
	"net/http"

	"github.com/lhj/backEnd/handlers"
	"github.com/lhj/backend/models"
)

func main() {
	debug := true

	models.InitDB("lhj.db")

	// The router used to handle user requests.
	router := http.NewServeMux()

	// The end points that have been implemented.
	handlers.CreateAllHandlers()

	// Apply the endpoints to the router.
	for _, handle := range handles.Handlers {
		router.HandleFunc(handle.Route, handle.Handler)
	}

	address := settings.GetSiteAddress(debug)
	timeOut := settings.GetTimeOutAmount(debug)

	// Create the server and set the timeout limits.
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: timeOut,
		ReadTimeout:  timeOut,
	}

	log.Fatal(srv.ListenAndServe())
}
