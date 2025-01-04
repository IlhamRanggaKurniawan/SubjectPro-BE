package entity

import "time"

type Subject struct {
	Id        uint64     `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"not null;uniqueIndex:idx_class_subject"`
	ClassId   uint64     `gorm:"not null;uniqueIndex:idx_class_subject"`
	Schedules []Schedule `gorm:"foreignKey:SubjectId;constraint:onDelete:CASCADE;"`
	Tasks     []Task     `gorm:"foreignKey:SubjectId;constraint:onDelete:CASCADE;"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
