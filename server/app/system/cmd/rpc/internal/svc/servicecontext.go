package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"server/app/system/cmd/rpc/internal/config"
	"server/app/system/model"
	"server/common"
	"server/common/xconfig"
)

type ServiceContext struct {
	Config config.Config

	MenuModel model.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 同步数据库
	err := initDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}
	db := newGormDB(c.Mysql)
	return &ServiceContext{
		Config:    c,
		MenuModel: model.NewMenuModel(db),
	}
}

func initDB(c xconfig.Mysql) error {
	init := common.NewInitMysqlDB(c.Username, c.Password, c.Path, c.Port)
	return init.CreateDatabase(c.Dbname)
}

func newGormDB(c xconfig.Mysql) *gorm.DB {
	mc := mysql.Config{
		DSN:                       c.Dsn(),
		SkipInitializeWithVersion: false,
		DefaultStringSize:         191,
	}

	db, err := gorm.Open(mysql.New(mc))
	if err != nil {
		logx.Error("[USER RPC MYSQL CONNECTION ERROR] : ", err)
		os.Exit(0)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	logx.Info("[USER RPC MYSQL CONNECTION SUCCESS]")
	// 同步数据库表
	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Menu{}, // 系统菜单表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
