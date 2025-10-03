package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
			return
		}

		user := User{UserID: id}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	case "POST":
		var user User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil || user.Name == "" {
			http.Error(w, `{"error": "invalid name"}`, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := map[string]string{"created": user.Name}
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
