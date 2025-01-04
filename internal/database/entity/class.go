package entity

import "time"

type Class struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Students  []User    `gorm:"foreignKey:ClassId;constraint:OnDelete:SET NULL"`
	Subjects  []Subject `gorm:"foreignKey:ClassId;constraint:onDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

