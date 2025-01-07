package task

import (

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(subjectId uint64, taskType string, note string, deadline string) (*entity.Task, error)
	FindAllByDeadline(deadline string, subjectId uint64) (*[]entity.Task, error)
	Delete(id uint64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(subjectId uint64, taskType string, note string, deadline string) (*entity.Task, error) {
	task := entity.Task{
		Type:      taskType,
		Note:      note,
		Deadline:  deadline,
		SubjectId: subjectId,
	}

	err := r.db.Save(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) FindAllByDeadline(deadline string, subjectId uint64) (*[]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Where("deadline = ? AND subject_id = ?", deadline, subjectId).Find(&tasks).Error

	if err != nil {
		return nil,err
	}

	return &tasks, nil
}

func (r *taskRepository) Delete(id uint64) error {
	err := r.db.Delete(entity.Task{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
