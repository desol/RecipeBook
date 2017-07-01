package handlers

import (
	"encoding/json"
	"net/http"

	"time"

	"strconv"

	"github.com/lhj/backend/models"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func createAccountHandlers() {
	handler.Route = "/auth"  // Set the path for this handler
	handler.Handler = auth   // Register which function will be triggered
	handler.Use(RecoverWrap) // Register which wrappers should be applied when the function is triggered
	Handlers = append(Handlers, handler)
}

func auth(w http.ResponseWriter, r *http.Request) {
	var id models.Identity
	switch r.Method {
	case "POST":
		un, pw, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Model State Not Valid.", http.StatusBadRequest)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to make hashbrows", http.StatusInternalServerError)
		}
		id, err = models.Login(un, hash)
		if err != nil {
			http.Error(w, "Invalid Credentials.", http.StatusUnauthorized)
		}

		jsonID, err := json.Marshal(id)
		if err != nil {
			http.Error(w, "Failed to produce identity", http.StatusInternalServerError)
		}
		w.Write(jsonID)

	case "PATCH":
		var err error
		header := r.Header
		id.Email = header.Get("email")
		id.DisplayName = header.Get("displayname")
		id.Token, err = uuid.FromBytes([]byte(header.Get("token")))
		if err != nil {
			http.Error(w, "Invalid Token.", http.StatusUnauthorized)
			return
		}
		id.Expires, err = time.Parse("", header.Get("expires"))
		if err != nil {
			http.Error(w, "Invalid Token Expiration.", http.StatusUnauthorized)
			return
		}
		err = models.CheckToken(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		header.Set("expires", id.Expires.String())

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
