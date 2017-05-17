package handlers

import (
	"errors"
	"log"
	"net/http"

	settings "github.com/RecipeBook/BackEnd/Settings"
)

// Use : Allows the user to pass which ever wrappers are necessary for a Handler and applies the wrappers to the Handler.
func (h *Handler) Use(middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	// Loop through each of the wrappers passed in.
	for _, m := range middleware {
		// Apply the wrapper to the handler.
		h.Handler = m(h.Handler)
	}
}

// RecoverWrap : Recovers if an endpoint panics and logs the error.
func RecoverWrap(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error // The error which caused the panic.

		// The function to catch the panic as the handler exits due to the error.
		defer func() {
			r := recover() // The object which will catch the panic

			// If there was a panic to recover from
			if r != nil {
				// Which type of panic was it?
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}

				log.Print(err)                                             // Log the error
				http.Error(w, err.Error(), http.StatusInternalServerError) // Notify the caller that an error occurred.
			}
		}()
		h.ServeHTTP(w, r)
	})
}

// AuthWrap : Checks the user's header for authentication prior to allowing the user to reach the endpoint.
func AuthWrap(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtain the credentials needed for this endpoint.
		correctUsername, correctPassword := settings.GetCredentials()
		// Obtain the user's credentials.
		username, password, _ := r.BasicAuth()

		// Check if the user's credentials and the endpoint's required credentials match.
		if username != correctUsername || password != correctPassword {
			// If the credentials don't match, then send an error to the caller.
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			// Log the fact a failed log in attempt happened.
			log.Printf("AUTHORIZATION FAILED~ Username: %s & Password: %s", username, password)
			return
		}
		h.ServeHTTP(w, r)
	})
}
