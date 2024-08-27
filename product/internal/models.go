package internal

type Product struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}
