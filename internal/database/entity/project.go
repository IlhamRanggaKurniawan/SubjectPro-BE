package entity

import "time"

type Project struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"not null"`
	Members   []User `gorm:"many2many:project_member;constraint:OnDelete:CASCADE;"`
	Status    string `gorm:"not null;default:development"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
