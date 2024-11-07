package entity

type RegisterRequest struct {
	Email    string
	Name     string
	Username string
	Password string
}

type LoginRequest struct {
	Username string
	Password string
}
