package entity

import "time"

type Team struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	Owner     string `gorm:"not null"`
	Admins    []User `gorm:"many2many:team_admin;constraint:onDelete:CASCADE;"`
	Members   []User `gorm:"many2many:team_member;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
