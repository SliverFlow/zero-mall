package model

import "gorm.io/gorm"

type (
	ProductModel interface {
		productModel
		categoryModel
	}

	customProductModel struct {
		*defaultProductModel
		*defaultCategoryModel
	}
)

func NewProductModel(db *gorm.DB) ProductModel {
	return customProductModel{
		defaultProductModel:  newDefaultProductModel(db),
		defaultCategoryModel: newDefaultCategoryModel(db),
	}
}
