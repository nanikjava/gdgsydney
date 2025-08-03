package routes

import (
	"encoding/json"
	"gdgsydney/db"
	"net/http"
)

type API struct {
	DB *db.Database
}

func (a API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Error processing request", http.StatusBadRequest)
		return
	}

	err = AuthenticateUser(a.DB.DB, req.Username, req.Password)

	if err != nil {
		http.Error(w, "Login failure", http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
