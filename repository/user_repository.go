package repository

import (
	"context"
	"grpc_hello_world/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	DeleteByID(ctx context.Context, id string) error
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
}
