package domain

// Bebas dari dependensi external (no echo, no gorm)
type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Mengatur nama tabel secara eksplisit agar GORM tidak menjamakkannya
func (Product) TableName() string {
	return "product" // atau "tbl_product" sesuai kebutuhan Anda
}