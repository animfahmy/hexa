package ports

import "hexa/internal/core/domain"

// Driven Port (Outbound): Apa yang dibutuhkan core dari luar (misal: database)
type UserRepository interface {
	GetData() domain.User
}

// Driving Port (Inbound): Apa yang bisa dilakukan core (Logika Bisnis/Use Case)
type UserService interface {
	GetProfile() string
}



// Driven Port (Outbound): Apa yang dibutuhkan Core dari database
type ProductRepository interface {
	FindAll() ([]domain.Product, error)
}

// Driving Port (Inbound): Apa yang bisa dilakukan oleh Core (Use Case)
type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
}