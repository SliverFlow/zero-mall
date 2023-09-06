package model

import "gorm.io/gorm"

type (
	UserModel interface {
		userModel
		businessModel
	}

	customUserModel struct {
		*defaultUserModel
		*defaultBusinessModel
	}
)

func NewUserModel(db *gorm.DB) UserModel {
	return &customUserModel{
		defaultUserModel:     newUserModel(db),
		defaultBusinessModel: newBusinessModel(db),
	}
}
