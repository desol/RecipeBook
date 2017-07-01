package handlers

import (
	"net/http"
)

// Handler : A handler has a route to handle and a function to be triggered when activated.
type Handler struct {
	Route   string
	Handler http.HandlerFunc
}

// Handlers : A slice of handlers which is used by routing to determine which paths can be 'handled' by the server
var Handlers []Handler

// A specific handler for an endpoint
var handler Handler

// CreateAllHandlers : A blanket function for when all handlers will be used.
func CreateAllHandlers() {

	// Mainly for testing and debugging, no real purpose
	createTestHandlers()

	createAccountHandlers()

}

// createTestHandlers : Registers all of the handlers that are associated to the CER extension.
func createTestHandlers() {
	handler.Route = "/tasks/testconnection" // Set the path for this handler
	handler.Handler = pingpong              // Register which function will be triggered
	handler.Use(RecoverWrap)                // Register which wrappers should be applied when the function is triggered
	Handlers = append(Handlers, handler)

	handler.Route = "/tasks/testauthentication" // Set the path for this handler
	handler.Handler = authorized                // Register which function will be triggered
	handler.Use(AuthWrap, RecoverWrap)          // Register which wrappers should be applied when the function is triggered
	Handlers = append(Handlers, handler)
}

// pingpong : Allows a user to determine if they have found where the Task System is being hosted from.
func pingpong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connected"))
}

// authorized : Allows a user to determine if their credentials work.
func authorized(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authorized"))
}
