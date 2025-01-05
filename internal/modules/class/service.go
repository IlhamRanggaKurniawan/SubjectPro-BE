package class

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/user"
)

type ClassService interface {
	CreateClass(userId uint64, name string) (*entity.Class, error)
	FindClass(id uint64) (*entity.Class, error)
	AddStudents(id uint64, newStudentsIds []uint64) (*entity.Class, error)
}

type classService struct {
	classRepo ClassRepository
	userRepo  user.UserRepository
}

func NewService(classRepo ClassRepository, userRepo user.UserRepository) ClassService {
	return &classService{classRepo: classRepo, userRepo: userRepo}
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

func (s *classService) AddStudents(id uint64, newStudentsIds []uint64) (*entity.Class, error) {
	users, err := s.userRepo.FindManyById(newStudentsIds)

	if err != nil {
		return nil, err
	}

	if len(*users) == 0 {
		return nil, fmt.Errorf("no users found for the given IDs: %v", newStudentsIds)
	}

	class, err := s.classRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	class.Students = append(class.Students, *users...)

	class, err = s.classRepo.Update(class)

	if err != nil {
		return nil, err
	}

	return class, nil
}
