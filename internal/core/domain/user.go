package domain

// Logika bisnis murni, tidak ada import "net/http" atau "database/sql"
type User struct {
	Name string
	Role string
}