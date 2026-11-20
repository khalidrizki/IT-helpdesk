package account

import (
	"context"
	"strings"
	"unicode/utf8"
)

var deptList = map[int]string{
	1:  "business-development",
	2:  "customer-service",
	3:  "finance-accounting",
	4:  "human-resources",
	5:  "it",
	6:  "legal-compliance",
	7:  "logistics",
	8:  "management",
	9:  "marketing",
	10: "operations",
	11: "procurement",
	12: "sales",
}

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateAccount(
	ctx context.Context,
	reqData CreateAccountInput,
) error {
	var validationErrors []string

	name := strings.TrimSpace(reqData.Name)
	if name == "" {
		validationErrors = append(validationErrors, "name is required")
	} else if utf8.RuneCountInString(name) < 3 {
		validationErrors = append(validationErrors, "name must contain at least 3 characters")
	}

	if _, exists := deptList[reqData.DepartmentID]; !exists {
		validationErrors = append(validationErrors, "invalid department")
	}

	// if len(validationErrors) > 0 {
	// 	return errors.New(strings.Join(validationErrors, "; "))
	// }
	err := s.repository.CreateAccount(ctx, CreateAccountInput{
		Name:         name,
		DepartmentID: reqData.DepartmentID,
	})
	if err != nil {
		return err
	}
	return nil
}
