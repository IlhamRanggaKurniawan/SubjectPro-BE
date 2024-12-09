package entity

import "time"

type User struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"default:member"`
	ClassId   *uint64   `gorm:"default:null"`
	Motto     *string   `gorm:"default:null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
