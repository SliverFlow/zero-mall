package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type (
	ProductModel interface {
		productModel
		categoryModel
		fkProductModel
	}

	customProductModel struct {
		*defaultProductModel
		*defaultCategoryModel
		*defaultFkProductModel
	}
)

func NewProductModel(db *gorm.DB, rdb *redis.Client) ProductModel {
	return customProductModel{
		defaultProductModel:   newDefaultProductModel(db, rdb),
		defaultCategoryModel:  newDefaultCategoryModel(db),
		defaultFkProductModel: newDefaultFkProductModel(db, rdb),
	}
}
