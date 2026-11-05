package account

import (
    "net/mail"
    "strings"
    "unicode/utf8"
    "errors"
)

type RegisterInput struct {
    Name            string
    Email           string
    DepartmentID    int
    Password        string
    ConfirmPassword string
}

var deptList = map[int]string{
	1: "business-development",
	2: "customer-service", 
	3: "finance-accounting", 
	4: "human-resources", 
	5: "it", 
	6: "legal-compliance", 
	7: "logistics", 
	8: "management", 
	9: "marketing", 
	10: "operations", 
	11: "procurement", 
	12: "sales",
}

type Service struct {
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) RegisterAccount(reqData RegisterInput) []string {   
    var validationErrors []string

    name := strings.TrimSpace(reqData.Name)

    if name == "" {
        validationErrors = append(validationErrors, "invalid name provided")
    }

    if utf8.RuneCountInString(name) < 3 {
        validationErrors = append(validationErrors, "name too short")
    }

    if strings.TrimSpace(reqData.Email) == "" {
        validationErrors = append(validationErrors, "email is required")
    }

    addr, err := mail.ParseAddress(reqData.Email)
    if err != nil || addr.Address != reqData.Email {
        validationErrors = append(validationErrors, "email format")
    }

    if _, exists := deptList[reqData.DepartmentID]; !exists {
        validationErrors = append(validationErrors, "invalid department")
    }

    if strings.TrimSpace(reqData.Password) == "" {
        validationErrors = append(validationErrors, "password is required")
    }

    if strings.TrimSpace(reqData.ConfirmPassword) == "" {
        validationErrors = append(validationErrors, "confirm password is required")
    }

    if reqData.Password != reqData.ConfirmPassword {
        validationErrors = append(validationErrors, "password and confirm password do not match")
    }

    if len(validationErrors) > 0 {
		http.Error(w, strings.Join(validationErrors, "; "), http.StatusBadRequest)
		return
	}
}

