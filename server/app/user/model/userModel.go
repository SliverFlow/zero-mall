package model

import "gorm.io/gorm"

type (
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func NewUserModel(db *gorm.DB) UserModel {
	return &customUserModel{
		newUserModel(db),
	}
}
