package common

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/common/xconfig"
)

type InitMysqlDB struct {
	port     int
	ip       string
	username string
	password string
}

// newInitMysqlDB
// Author [SliverFlow]
// @desc 实例化
func newInitMysqlDB(username, password, ip string, port int) *InitMysqlDB {
	return &InitMysqlDB{
		username: username,
		password: password,
		ip:       ip,
		port:     port,
	}
}

// dsn
// Author [SliverFlow]
// @desc 获取 dsn
func (im *InitMysqlDB) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/", im.username, im.password, im.ip, im.port)
}

// createDatabase
// Author [SliverFlow]
// @desc 创建数据库
func (im *InitMysqlDB) createDatabase(dbName string) error {
	db, err := sql.Open("mysql", im.dsn())
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logx.Error(err.Error())
		}
	}(db)
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbName))
	return err
}

func InitDB(c xconfig.Mysql) (db *gorm.DB, err error) {
	mysqlDB := newInitMysqlDB(c.Username, c.Password, c.Path, c.Port)
	err = mysqlDB.createDatabase(c.Dbname)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		return nil, err
	}
	mc := mysql.Config{
		DSN:                       c.Dsn(),
		SkipInitializeWithVersion: false,
		DefaultStringSize:         191,
	}
	db, err = gorm.Open(mysql.New(mc))
	if err != nil {
		logx.Error("[MYSQL CONNECTION ERROR] : ", err)
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	logx.Info("[ MYSQL CONNECTION SUCCESS]")
	return db, nil
}
