package main

import (
	"encoding/json"
	"net/http"

	"github.com/khalidrizki/IT-helpdesk/back-end/auth-service/internal/account"
)

func (s *server) handleRegister(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // body dibatasi max. 1 MiB
	var reqData account.RegisterInput

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := s.accountService.RegisterAccount(r.Context(), reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "account successfully created",
	})
}
