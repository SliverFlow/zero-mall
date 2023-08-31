package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"server/app/user/cmd/rpc/internal/config"
	"server/app/user/model"
	"server/common/xconfig"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := newGormDB(c.Mysql)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(db),
	}
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
		&model.User{}, // 系统菜单表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
