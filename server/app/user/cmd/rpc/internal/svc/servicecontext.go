package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/user/cmd/rpc/internal/config"
	"server/app/user/model/auto"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := newGormDB(c.Mysql)
	// 同步数据库
	err := db.AutoMigrate(&auto.User{})
	if err != nil {
		logx.Error("[USER RPC MYSQL AutoMigrate ERROR] : ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[USER RPC MYSQL AutoMigrate SUCCESS]")

	return &ServiceContext{
		Config: c,
	}
}

func newGormDB(c gormc.Mysql) *gorm.DB {
	db, err := gormc.ConnectMysql(c)
	if err != nil {
		logx.Error("[USER RPC MYSQL CONNECTION ERROR] : ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[USER RPC MYSQL CONNECTION SUCCESS]")
	return db
}
