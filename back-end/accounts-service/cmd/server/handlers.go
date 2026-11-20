package main

import (
	"encoding/json"
	"net/http"

	"github.com/khalidrizki/IT-helpdesk/back-end/accounts-service/internal/account"
)

type server struct {
	accountService *account.Service
}

func (s *server) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input account.CreateAccountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := s.accountService.CreateAccount(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "account successfully created",
	})
}
