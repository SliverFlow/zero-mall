package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/cart/cmd/rpc/internal/config"
	"server/app/cart/model"
	"server/common"
)

type ServiceContext struct {
	Config config.Config

	CartModel model.CartModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := common.InitDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}
	autoMigrate(db)

	return &ServiceContext{
		Config:    c,
		CartModel: model.NewCartModel(db),
	}
}

// 同步表
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Cart{}, // 购物车表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
