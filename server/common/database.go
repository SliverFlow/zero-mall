package common

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitMysqlDB struct {
	port     int
	ip       string
	username string
	password string
}

// NewInitMysqlDB
// Author [SliverFlow]
// @desc 实例化
func NewInitMysqlDB(username, password, ip string, port int) *InitMysqlDB {
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

// CreateDatabase
// Author [SliverFlow]
// @desc 创建数据库
func (im *InitMysqlDB) CreateDatabase(dbName string) error {
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
