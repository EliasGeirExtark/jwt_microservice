package models

import (
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	ExpireAt     time.Time `gorm:"column:expire_at; not null"`
	RefreshToken string    `gorm:"column:refresh_token; not null; primary key"`
	AccessToken  string    `gorm:"column:access_token; not null"`
}

// CreateRefresh generate or update the refresh token inside the database
func (r *RefreshToken) CreateRefresh(db *gorm.DB) error {
	var refresh RefreshToken
	if db.Model(&RefreshToken{}).Where("refresh_token = ?", r.RefreshToken).First(&refresh); refresh.RefreshToken == "" {
		refresh.ExpireAt = r.ExpireAt
		refresh.RefreshToken = r.RefreshToken
		refresh.AccessToken = r.AccessToken
		err := db.Create(&refresh).Error
		return err
	} else {
		refresh.ExpireAt = r.ExpireAt
		err := db.Save(&refresh).Error
		return err
	}
}

// IsRefreshTokenValid check if exists a refresh token inside the database and if it is not expired
func (r *RefreshToken) IsRefreshTokenValid(db *gorm.DB) (bool, error) {
	var refresh RefreshToken
	if err := db.Model(&RefreshToken{}).Where("refresh_token = ?", r.RefreshToken).Find(&refresh).Error; err != nil {
		return false, err
	} else {
		if refresh.ExpireAt.After(time.Now()) && refresh.AccessToken == r.AccessToken {
			return true, nil
		}
	}

	return false, nil
}

// DeleteRefreshToken deletes the current token from the database
func (r *RefreshToken) DeleteRefreshToken(db *gorm.DB) error {
	var refresh RefreshToken
	if err := db.Model(&RefreshToken{}).Where("refresh_token = ?", r.RefreshToken).First(&refresh).Error; err != nil {
		return err
	}

	return db.Where("refresh_token = ?", r.RefreshToken).Delete(&refresh).Error
}
