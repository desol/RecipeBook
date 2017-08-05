package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/lhj/backend/handlers"
	"github.com/lhj/backend/models"
	"github.com/lhj/backend/settings"
)

func main() {
	debug := true

	// Init the application's settings.
	settings.Populate(debug)

	// Init the application's storm DB
	models.InitDB(settings.Settings.StormDB)

	// The router used to handle user requests.
	router := http.NewServeMux()

	// The end points that have been implemented.
	handlers.CreateAllHandlers()

	// Apply the endpoints to the router.
	for _, handle := range handlers.Handlers {
		fmt.Println(handle.Route)
		router.HandleFunc(handle.Route, handle.Handler)
	}

	// Create the server and set the timeout limits.
	srv := &http.Server{
		Handler:      router,
		Addr:         settings.Settings.Port,
		WriteTimeout: settings.Settings.ServerTimeout,
		ReadTimeout:  settings.Settings.ServerTimeout,
		IdleTimeout:  settings.Settings.ServerIdleTimeout,
	}

	fmt.Println(srv.Addr)

	// Wait to serve
	log.Fatal(srv.ListenAndServe())
}
