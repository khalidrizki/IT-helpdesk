package account

type Account struct {
    ID           string
    Name         string
    Email        string
    Department   string
    PasswordHash string
    Role         string
    IsActive     bool
}