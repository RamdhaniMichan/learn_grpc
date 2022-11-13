package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID       string `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	Name       string `json:"name"`
	Salutation string `json:"salutation"`
}

func (u User) ValidateSave() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(5, 100)),
		validation.Field(&u.Salutation, validation.Required, validation.Length(10, 50)),
	)
}

func (u User) ValidateFind() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UUID, validation.Required),
	)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}

	return nil
}
