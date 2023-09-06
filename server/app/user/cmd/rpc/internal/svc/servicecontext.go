package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/user/cmd/rpc/internal/config"
	"server/app/user/model"
	"server/common"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
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
		UserModel: model.NewUserModel(db),
	}
}

// 同步表
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},     // 用户表
		&model.Business{}, // 商家表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
