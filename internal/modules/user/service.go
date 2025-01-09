package user

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type UserService interface {
	Register(username string, email string, password string) (*entity.User, error)
	Login(email string, password string) (*entity.User, error)
	FindUserLikeEmail(email string) (*[]entity.User, error)
	Update(id uint64, username string, email string, password string) (*entity.User, error)
}

type userService struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Register(username string, email string, password string) (*entity.User, error) {
	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.Create(username, email, hashedPassword)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email string, password string) (*entity.User, error) {
	user, err := s.userRepository.FindOneByEmail(email)

	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(user.Password, password)

	if err != nil {
		return nil, fmt.Errorf("the provided credentials do not match our records")
	}

	return user, nil
}

func (s *userService) FindUserLikeEmail(email string) (*[]entity.User, error) {
	users, err := s.userRepository.FindManyLikeEmail(email)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) Update(id uint64, username string, email string, password string) (*entity.User, error) {
	user, err := s.userRepository.FindOneById(id)

	if err != nil {
		return nil, err
	}

	if username != "" {
		user.Username = username
	}

	if email != "" {
		user.Email = email
	}

	if password != "" {
		hashedPassword, err := utils.HashPassword(password)

		if err != nil {
			return nil, err
		}

		user.Password = hashedPassword
	}

	user, err = s.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
