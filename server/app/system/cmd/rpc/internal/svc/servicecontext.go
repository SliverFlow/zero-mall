package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/system/cmd/rpc/internal/config"
	"server/app/system/model"
	"server/common"
)

type ServiceContext struct {
	Config config.Config

	SystemModel model.SystemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := common.InitDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}
	// autoMigrate(db)

	return &ServiceContext{
		Config:      c,
		SystemModel: model.NewSystemModel(db),
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Menu{}, // 系统菜单表
		&model.Role{}, // 系统角色表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
