package class

import (
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"gorm.io/gorm"
)

type ClassRepository interface {
	Create(userId uint64, name string) (*entity.Class, error)
	FindById(id uint64) (*entity.Class, error)
	Update(class *entity.Class) (*entity.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Create(userId uint64, name string) (*entity.Class, error) {
	admin := []entity.User{
		{
			Id: userId,
		},
	}

	class := entity.Class{
		Name:     name,
		Students: admin,
	}

	err := r.db.Save(&class).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Model(&entity.User{}).Where("id = ? AND role = ?", userId, "Student").Update("Role", "Class Leader").Error

	if err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *classRepository) FindById(id uint64) (*entity.Class, error) {
	var class entity.Class

	err := r.db.Preload("Students").Preload("Subjects").Where("id = ?", id).Take(&class).Error

	if err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *classRepository) Update(class *entity.Class) (*entity.Class, error) {
	err := r.db.Save(class).Error

	if err != nil {
		return nil, err
	}

	return class, err
}
