package subject

import (
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"gorm.io/gorm"
)

type SubjectRepository interface {
	Create(name string, ClassId uint64) (*entity.Subject, error)
	FindAllByClassId(classId uint64) (*[]entity.Subject, error)
	Delete(id uint64) error
}

type subjectRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) Create(name string, ClassId uint64) (*entity.Subject, error) {
	var subject = entity.Subject{
		Name:    name,
		ClassId: ClassId,
	}

	err := r.db.Save(&subject).Error

	if err != nil {
		return nil, err
	}

	return &subject, err
}

func (r *subjectRepository) FindAllByClassId(classId uint64) (*[]entity.Subject, error) {
	var subject []entity.Subject

	err := r.db.Where("class_id = ?", classId).Find(&subject).Error

	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *subjectRepository) Delete(id uint64) error {
	err := r.db.Delete(entity.Subject{}, id).Error

	return err
}
