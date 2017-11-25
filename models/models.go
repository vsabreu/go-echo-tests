package models

const (
	// AppUser defines user access
	AppUser Role = iota + 1

	// AppAdmin defines admin access
	AppAdmin
)

type (
	// Role defines access level
	Role int

	// User represents an application user
	User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Role   Role   `json:"role"`
		Active bool   `json:"active"`
	}

	// Product represents a product for sell
	Product struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"desc"`
		Price       float64 `json:"price"`
	}
)
