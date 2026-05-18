package domain

// Bebas dari dependensi external (no echo, no gorm)
type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}