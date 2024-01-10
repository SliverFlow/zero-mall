package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/log/cmd/rpc/internal/config"
	"server/app/log/model"
	"server/common"
)

type ServiceContext struct {
	Config   config.Config
	LogModel model.LogModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := common.InitDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}

	// autoMigrate(db)

	return &ServiceContext{
		Config:   c,
		LogModel: model.NewLogModel(db),
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Log{}, // 系统日志表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
