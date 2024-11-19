package entity

import "time"

type User struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:member"`
	Teams     []Team `gorm:"many2many:team_member;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
