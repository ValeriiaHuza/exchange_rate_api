package models

type User struct {
	ID    uint
	Email string `gorm:"unique;not null"`
}
