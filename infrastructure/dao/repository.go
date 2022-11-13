package dao

import (
	"grpc_hello_world/infrastructure/persistent"
	"grpc_hello_world/repository"

	"gorm.io/gorm"
)

type Repository struct {
	DB       *gorm.DB
	UserRepo repository.UserRepository
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB:       db,
		UserRepo: persistent.NewUserRepo(db),
	}
}
