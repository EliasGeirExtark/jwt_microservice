package models

import (
	"github.com/extark/jwt_microservice/utils"
	"time"
)

type RefreshToken struct {
	ExpireAt     time.Time `gorm:"column:expire_at; not null"`
	RefreshToken string    `gorm:"column:refresh_token; not null; primary key"`
	AccessToken  string    `gorm:"column:access_token; not null"`
}

// CreateRefresh generate or update the refresh token inside the database
func (r *RefreshToken) CreateRefresh() error {
	var refresh RefreshToken
	if utils.Cfg.SQLDB.Model(&RefreshToken{}).Where("refresh_token = ?", r.RefreshToken).First(&refresh); refresh.RefreshToken == "" {
		refresh.ExpireAt = r.ExpireAt
		refresh.RefreshToken = r.RefreshToken
		refresh.AccessToken = r.AccessToken
		err := utils.Cfg.SQLDB.Create(&refresh).Error
		return err
	} else {
		refresh.ExpireAt = r.ExpireAt
		err := utils.Cfg.SQLDB.Save(&refresh).Error
		return err
	}
}

// IsRefreshTokenValid check if exists a refresh token inside the database and if it is not expired
func (r *RefreshToken) IsRefreshTokenValid() (bool, error) {
	var refresh RefreshToken
	if err := utils.Cfg.SQLDB.Model(&RefreshToken{}).Where("refresh_token = ?", r.RefreshToken).Find(&refresh).Error; err != nil {
		return false, err
	} else {
		if refresh.ExpireAt.After(time.Now()) && refresh.AccessToken == r.AccessToken {
			return true, nil
		}
	}

	return false, nil
}

// DeleteRefreshToken deletes the current token from the database
func (r *RefreshToken) DeleteRefreshToken() error {
	var refresh RefreshToken
	if err := utils.Cfg.SQLDB.Model(&RefreshToken{}).Where("refresh_token = ?", r.RefreshToken).First(&refresh).Error; err != nil {
		return err
	}

	return utils.Cfg.SQLDB.Where("refresh_token = ?", r.RefreshToken).Delete(&refresh).Error
}
