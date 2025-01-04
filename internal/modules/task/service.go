package task

import (
	"time"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
)

type TaskService interface {
	CreateTask(subjectId uint64, taskType string, note string, deadline time.Time) (*entity.Task, error)
	DeleteTask(id uint64) error
}

type taskService struct {
	taskRepository TaskRepository
}

func NewService(taskRepository TaskRepository) TaskService {
	return &taskService{taskRepository: taskRepository}
}

func (s *taskService) CreateTask(subjectId uint64, taskType string, note string, deadline time.Time) (*entity.Task, error) {
	task, err := s.taskRepository.Create(subjectId, taskType, note, deadline)

	if err !=  nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id uint64) error {
	err := s.taskRepository.Delete(id)

	return err
}
