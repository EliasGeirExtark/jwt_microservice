package models

type Account struct {
	Nickname   string  `gorm:"column:name; not null"`
	Email      string  `gorm:"column:email; not null; unique"`
	Phone      string  `gorm:"column:phone; not null; unique"`
	ResetToken *string `gorm:"column:reset_token;"`
}
