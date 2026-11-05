package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/mail"
	"strings"
	"unicode/utf8"
	"back-end/auth-service/internal/account"
)

func (s *server) handleRegister(w http.ResponseWriter, r *http.Request){
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)  // body dibatasi max. 1 MiB
	byteData, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error reading request body", "err", err)
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	if len(byteData) == 0 {
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	var reqData account.RegisterInput
	err = json.Unmarshal(byteData, &reqData)
	if err != nil {
		slog.Error("error unmarshalling request body", "err", err)
		http.Error(w, "error parsing request JSON", http.StatusBadRequest)
		return
	}

	err := s.accountService.RegisterAccount(reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var output bytes.Buffer
	output.WriteString("Accout with name: '")
	output.WriteString(reqData.Name)
	output.WriteString("', email: '")
	output.WriteString(reqData.Email)
	output.WriteString("', and department: '")
	output.WriteString(deptList[reqData.DepartmentID])
	output.WriteString("' successfully created.\n")

	_, err = w.Write(output.Bytes())
	if err != nil {
		slog.Error("error writing response body", "err", err)
		return
	}
}