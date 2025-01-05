package user

import (
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(username string, email string, password string) (*entity.User, error)
	FindOneByEmail(email string) (*entity.User, error)
	FindOneById(id uint64) (*entity.User, error)
	FindManyById(ids []uint64) (*[]entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(username string, email string, password string) (*entity.User, error) {
	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneById(id uint64) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindManyById(ids []uint64) (*[]entity.User, error) {
	var users []entity.User

	err := r.db.Where("id IN ?", ids).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) Update(user *entity.User) (*entity.User, error) {
	err := r.db.Save(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
