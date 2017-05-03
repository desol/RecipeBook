package handlers

import (
	"net/http"
)

// Handlers : The splice of handles the task system will be using.
type Handlers struct {
	Handles []Handler
}

// Handler : A handler has a route to handle and a function to be triggered when activated.
type Handler struct {
	Route   string
	Handler http.HandlerFunc
}

// AddHandler : Extension for the Handlers struct which allows a handler to be added to the splice.
func (handles *Handlers) AddHandler(handle Handler) {
	handles.Handles = append(handles.Handles, handle)
}

// CreateAllHandlers : A blanket function for when all handlers will be used.
func CreateAllHandlers() Handlers {
	var allHandlers Handlers

	// Create a splice of all the general purpose handlers
	utilHandlers := CreateUtilHandlers()
	// Add the general purpose handles to the conglomerate of handlers
	for _, handle := range utilHandlers.Handles {
		allHandlers.AddHandler(handle)
	}

	return allHandlers // Return all the Task System's handlers
}

// CreateUtilHandlers : Registers all of the handlers that are associated to the CER extension.
func CreateUtilHandlers() Handlers {
	var handle Handler   // A specific handler for an endpoint
	var handles Handlers // All of the handlers for CERs

	handle.Route = "/tasks/testconnection" // Set the path for this handler
	handle.Handler = pingpong              // Register which function will be triggered
	handle.Use(RecoverWrap)                // Register which wrappers should be applied when the function is triggered
	handles.AddHandler(handle)             // Add the handle to the splice of handles for the CER extension

	handle.Route = "/tasks/testauthentication" // Set the path for this handler
	handle.Handler = authorized                // Register which function will be triggered
	handle.Use(AuthWrap, RecoverWrap)          // Register which wrappers should be applied when the function is triggered
	handles.AddHandler(handle)                 // Add the handle to the splice of handles for the CER extension

	return handles // Return the general purpose handles for the task system
}

// pingpong : Allows a user to determine if they have found where the Task System is being hosted from.
func pingpong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connected"))
}

// authorized : Allows a user to determine if their credentials work.
func authorized(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authorized"))
}
