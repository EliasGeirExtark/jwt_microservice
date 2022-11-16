package models

import (
	"github.com/extark/jwt_microservice/utils"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Model
	Nickname   string  `gorm:"column:nickname;check:nickname is not null or email is not null or phone is not null"`
	Email      string  `gorm:"column:email;check:nickname is not null or email is not null or phone is not null"`
	Phone      string  `gorm:"column:phone;check:nickname is not null or email is not null or phone is not null"`
	Password   string  `gorm:"column:password"`
	ResetToken *string `gorm:"column:reset_token;check:nickname is not null or email is not null or phone is not null"`
}

// CheckLogin checks if the password sent is correct and returns a user and an eventual error
// userTypeID a string type, it defines which is the primary key
func CheckLogin(userTypeID string, username string, password string) (*Account, error) {
	// Get the user by the user type id
	var account Account
	if err := utils.Cfg.SQLDB.Where(userTypeID+" = ?", username).Take(&account).Error; err != nil {
		return nil, err
	}

	// Compare the hash on db and the passed password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return &account, nil
}
