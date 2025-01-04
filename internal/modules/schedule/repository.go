package schedule

import (
	"time"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(day string, subjectId uint64, startTime time.Time, endTime time.Time) (*entity.Schedule, error)
	Delete(id uint64) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(day string, subjectId uint64, startTime time.Time, endTime time.Time) (*entity.Schedule, error) {
	schedule := entity.Schedule{
		Day:       day,
		SubjectId: subjectId,
		StartTime: startTime,
		EndTime:   endTime,
	}

	err := r.db.Save(&schedule).Error

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (r *scheduleRepository) Delete(id uint64) error {
	err := r.db.Delete(entity.Schedule{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
