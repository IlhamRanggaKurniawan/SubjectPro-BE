package entity

import "time"

type Class struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Admin     []User    `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE;"`
	Students  []User    `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE;"`
	Subjects  []Subject `gorm:"foreignKey:ClassId;constraint:onDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
