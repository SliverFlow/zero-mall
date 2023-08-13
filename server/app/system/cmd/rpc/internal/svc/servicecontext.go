package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/system/cmd/rpc/internal/config"
	"server/app/system/model/auto"
	"server/common"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 同步数据库
	err := initDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}
	// 同步表
	db := newGormDB(c.Mysql)
	if err = db.AutoMigrate(&auto.Menu{}); err != nil {
		logx.Error("AutoMigrate mysql table err", err.Error())
		os.Exit(0)
	}
	return &ServiceContext{
		Config: c,
	}
}

func initDB(c gormc.Mysql) error {
	init := common.NewInitMysqlDB(c.Username, c.Password, c.Path, c.Port)
	return init.CreateDatabase(c.Dbname)
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
