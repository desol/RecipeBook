package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/lhj/BackEnd/Handlers"
)

func main() {
	debug := true

	// The router used to handle user requests.
	router := mux.NewRouter()

	// The end points that have been implemented.
	handles := handlers.CreateAllHandlers()

	// Apply the endpoints to the router.
	for _, handle := range handles.Handles {
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
