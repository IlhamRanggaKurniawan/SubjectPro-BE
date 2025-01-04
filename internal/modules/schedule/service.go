package schedule

import (
	"time"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
)

type ScheduleService interface {
	CreateSchedule(day string, subjectId uint64, startTime time.Time, endTime time.Time) (*entity.Schedule, error)
	DeleteSchedule(id uint64) error
}

type scheduleService struct {
	scheduleRepository ScheduleRepository
}

func NewService(scheduleRepository ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepository: scheduleRepository}
}

func (s *scheduleService) CreateSchedule(day string, subjectId uint64, startTime time.Time, endTime time.Time) (*entity.Schedule, error) {
	schedule, err := s.scheduleRepository.Create(day, subjectId, startTime, endTime)

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (s *scheduleService) DeleteSchedule(id uint64) error {
	err := s.scheduleRepository.Delete(id)

	return err
}
