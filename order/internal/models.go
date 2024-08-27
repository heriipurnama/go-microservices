package internal

type Order struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProductID string `gorm:"not null"`
	Quantity  int32  `gorm:"not null"`
	UserID    string `gorm:"not null"`
}
