package persistent

import (
	"context"
	"grpc_hello_world/entity"
	"grpc_hello_world/repository"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u UserRepo) GetAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := u.db.Find(users).Error
	return users, err
}

func (u UserRepo) Create(ctx context.Context, user *entity.User) error {
	err := u.db.Create(&user).Error
	return err
}

func (u UserRepo) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("uuid = ?", id).First(&user).Error

	return &user, err
}

func (u UserRepo) DeleteByID(ctx context.Context, id string) error {
	var user entity.User
	err := u.db.Where("uuid = ?", id).Delete(&user).Error
	return err
}

func (u UserRepo) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	var userUpdate entity.User

	err := u.db.Model(&userUpdate).Where("uuid = ?", userUpdate.UUID).Updates(map[string]interface{}{"name": user.Name, "salutation": user.Salutation}).Error

	return &userUpdate, err

}

var _ repository.UserRepository = &UserRepo{}
