package model

import "gorm.io/gorm"

type (
	UserModelGorm interface {
		userModelGorm
	}

	customUserModelGorm struct {
		*defaultUserModelGorm
	}
)

func NewUserModelGorm(db *gorm.DB) UserModelGorm {
	return &customUserModelGorm{
		newUserModelGorm(db),
	}
}
