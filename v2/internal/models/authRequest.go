package models

// The request Dto for both register and login
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
