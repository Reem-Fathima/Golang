package models

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Priority    string
	Deadline    string
	Status      string
	Category    string
	UserID      uint
}
