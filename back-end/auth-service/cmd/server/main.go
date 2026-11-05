import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/mail"
	"strings"
	"unicode/utf8"
	"back-end/auth-service/internal/account"
)

type server struct {
    accountService *account.Service
}

func main(){
	accountSvc := account.NewService()

	s := &server{
		accountService: accountSvc,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", handleRegister)

	log.Fatal(http.ListenAndServe(":8080", mux))
}