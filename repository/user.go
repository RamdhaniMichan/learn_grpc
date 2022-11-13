package repository

import (
	"fmt"
	"grpc_hello_world/entity"

	"gorm.io/gorm"
)

func SaveUser(db *gorm.DB, User *entity.User) (*entity.User, error) {
	err := db.Create(User).Error
	fmt.Printf("User: %v\n", User)
	if err != nil {
		return nil, err
	}

	return User, nil
}
