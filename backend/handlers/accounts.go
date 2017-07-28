package handlers

import (
	"encoding/json"
	"net/http"

	"context"
	"strconv"
	"strings"
	"time"

	"../models"
)

func createAccountHandlers() {
	handler.Route = "/api/auth"        // Set the path for this handler
	handler.Handler = auth             // Register which function will be triggered
	handler.Use(recoverWrap, authWrap) // Register which wrappers should be applied when the function is triggered
	Handlers = append(Handlers, handler)

	handler.Route = "/api/account"        // Set the path for this handler
	handler.Handler = account             // Register which function will be triggered
	handler.Use(recoverWrap, accountWrap) // Register which wrappers should be applied when the function is triggered
	Handlers = append(Handlers, handler)
}

func auth(w http.ResponseWriter, r *http.Request) {
	var id models.Identity // To store basic user info.

	switch r.Method {
	// Logs the user in and returns a token
	case "POST":
		un, pw, ok := r.BasicAuth() // Get the basic auth request info.
		if !ok {
			// Return the unexpected error.
			http.Error(w, "Model State Not Valid.", http.StatusBadRequest)
			return
		}
		// Log the user in using the username and password from the basic auth
		id, err := models.Login(strings.ToLower(strings.TrimSpace(un)), strings.TrimSpace(pw))
		if err != nil {
			// Either bad credentials or the user was locked out or something unexpected
			http.Error(w, "Failed to validate credentials. Message: "+err.Error(), http.StatusUnauthorized)
		}

		// Return the user's identity
		jsonID, err := json.Marshal(id)
		if err != nil {
			http.Error(w, "Failed to produce identity", http.StatusInternalServerError)
		}
		w.Write(jsonID)

	// Updates and checks user's token
	case "PATCH":
		var err error
		header := r.Header
		id.Pk, err = strconv.Atoi(header.Get("pk"))
		if err != nil {
			http.Error(w, "Failed to read in PK", http.StatusInternalServerError)
		}
		id.Email = header.Get("email")
		id.DisplayName = header.Get("displayname")
		id.Token = header.Get("token")
		if err != nil {
			http.Error(w, "Invalid Token. Message: "+err.Error(), http.StatusUnauthorized)
			return
		}
		id.Expires, err = strconv.ParseInt(header.Get("expires"), 10, 64)
		if err != nil {
			http.Error(w, "Invalid Token Expiration: "+err.Error(), http.StatusUnauthorized)
			return
		}
		err = models.CheckToken(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		header.Set("expires", string(id.Expires))

	// Add a permission to the user.
	case "PUT":
		var perm models.Permission
		var err error
		perm.Active, err = strconv.ParseBool(r.Form.Get("active"))
		if err != nil {
			http.Error(w, "Invalid active value.", http.StatusNotAcceptable)
			return
		}
		perm.Pk, err = strconv.Atoi(r.Form.Get("pk"))
		if err != nil {
			http.Error(w, "Invalid pk.", http.StatusNotAcceptable)
			return
		}
		perm.Name = r.Form.Get("name")
		err = models.UpdatePermission(&perm)
		if err != nil {
			http.Error(w, "Failed to update permission", http.StatusInternalServerError)
			return
		}
		jsonReturn, err := json.Marshal(true)
		if err != nil {
			http.Error(w, "Failed to make a json bool...", http.StatusInternalServerError)
			return
		}
		w.Write(jsonReturn)

	// Returns if the user is permitted to continue or not
	case "GET":
		permitted := false
		perm, err := strconv.Atoi(r.URL.Query().Get("permission"))
		if err != nil {
			http.Error(w, "Invalid Permission.", http.StatusBadRequest)
		}
		err = models.CheckPermission(r.Header.Get("email"), perm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		ok, err := json.Marshal(permitted)
		if err != nil {
			http.Error(w, "Failed to produce boolean.", http.StatusInternalServerError)
		}
		w.Write(ok)
	}
}

// Functions dealing with a user's account
func account(w http.ResponseWriter, r *http.Request) {
	// Which type of request came in?
	switch r.Method {
	// User to create an account
	case "POST":
		id := new(models.Identity) // To store basic user info.
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(id) // Get the requested password
		if err != nil {
			http.Error(w, "Error Decoding ID. Message: "+err.Error(), http.StatusInternalServerError)
		}
		err = models.RegisterAccount(id) // Registed the info as a new account
		if err != nil {
			// Return the error info and don't allow the user to continue
			http.Error(w, "An unexpected error occurred while creating your account: "+err.Error(), http.StatusInternalServerError)
			return
		}
		id.Password = ""

		// Return the user's new identity
		jsonID, err := json.Marshal(id)
		if err != nil {
			http.Error(w, "Failed to produce identity", http.StatusInternalServerError)
		}
		w.Write(jsonID)

	case "PATCH":
		user := new(models.UserInfo)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(user)      // Get the requested password
		err = models.UpdateAccount(user) // Registed the info as a new account
		if err != nil {
			// Return the error info and don't allow the user to continue
			http.Error(w, "An unexpected error occurred while updating the account: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// Notify that the update succeeded
		jsonResponse, err := json.Marshal(true)
		if err != nil {
			http.Error(w, "Failed to make a jsonified boolean...", http.StatusInternalServerError)
		}
		w.Write(jsonResponse)

	}
}

func accountWrap(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutInSeconds := 120
		switch r.Method {
		case "PATCH":
			id := new(models.Identity)
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(id) // Get the requested password
			if err != nil {
				http.Error(w, "Error Decoding ID. Message: "+err.Error(), http.StatusInternalServerError)
				return
			}
			err = models.CheckToken(id)
			if err != nil {
				http.Error(w, "Error verifying token: "+err.Error(), http.StatusForbidden)
				return
			}
			err = models.CheckPermission(id.Email, 0)
			if err != nil {
				http.Error(w, "Not Permitted: "+err.Error(), http.StatusUnauthorized)
				return
			}
		}
		ctx, ctxCancel := context.WithTimeout(r.Context(), time.Duration(timeoutInSeconds)*time.Second)
		defer ctxCancel()
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authWrap(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutInSeconds := 60
		switch r.Method {
		case "PUT":
			id := new(models.Identity)
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(id) // Get the requested password
			if err != nil {
				http.Error(w, "Error Decoding ID. Message: "+err.Error(), http.StatusInternalServerError)
				return
			}
			err = models.CheckToken(id)
			if err != nil {
				http.Error(w, "Error verifying token: "+err.Error(), http.StatusForbidden)
				return
			}
			err = models.CheckPermission(id.Email, 0)
			if err != nil {
				http.Error(w, "Not Permitted: "+err.Error(), http.StatusUnauthorized)
				return
			}
		}
		ctx, ctxCancel := context.WithTimeout(r.Context(), time.Duration(timeoutInSeconds)*time.Second)
		defer ctxCancel()
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
