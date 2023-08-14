package service

import (
	"errors"
	"gorm.io/gorm"
	"server/app/system/model/auto"
	"server/common/xerr"
)

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{
		db: db,
	}
}

func (s *MenuService) Find(id uint) (*auto.Menu, error) {
	var enter auto.Menu
	if err := s.db.Model(&auto.Menu{}).Where("id = ?", id).First(&enter).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerr.NewCodeError(100001)
		}
		return nil, err
	}
	return &enter, nil
}
