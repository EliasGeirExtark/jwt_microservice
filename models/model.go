package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID      string         `gorm:"type:uuid;primary_key;column:uuid;default:uuid_generate_v4()" json:"uuid"`
	CreatedAt time.Time      `gorm:"column:created_at; <-:create; default:current_timestamp; not null" json:"-"`
	UpdateAt  time.Time      `gorm:"column:updated_at; default:current_timestamp; check: (updated_at >= created_at) OR updated_at is null" json:"-"`
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at; check: (deleted_at >= created_at AND deleted_at >= updated_at) OR deleted_at is null" json:"-"`
}
