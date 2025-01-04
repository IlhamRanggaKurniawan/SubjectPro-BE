package class

import (
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
)

type ClassService interface {
	CreateClass(userId uint64, name string) (*entity.Class, error)
	FindClass(id uint64) (*entity.Class, error)
	AddStudents(id uint64, newStudents []entity.User) (*entity.Class, error)
}

type classService struct {
	classRepo ClassRepository
}

func NewService(classRepo ClassRepository) ClassService {
	return &classService{classRepo: classRepo}
}

func (s *classService) CreateClass(userId uint64, name string) (*entity.Class, error) {
	class, err := s.classRepo.Create(userId, name)

	if err != nil {
		return nil, err
	}

	return class, nil
}

func (s *classService) FindClass(id uint64) (*entity.Class, error) {
	class, err := s.classRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	return class, nil
}

func (s *classService) AddStudents(id uint64, newStudents []entity.User) (*entity.Class, error) {
	class, err := s.classRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	class.Students = append(class.Students, newStudents...)

	class, err = s.classRepo.Update(class)

	if err != nil {
		return nil, err
	}

	return class, nil
}
