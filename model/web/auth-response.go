package web

import "time"

type AuthResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}
