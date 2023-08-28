package xconfig

import "fmt"

// Mysql gorm connection mysql config
type Mysql struct {
	Path          string // 服务器地址
	Port          int    `json:",default=3306"`                                               // 端口
	Config        string `json:",default=charset%3Dutf8mb4%26parseTime%3Dtrue%26loc%3DLocal"` // 高级配置
	Dbname        string // 数据库名
	Username      string // 数据库用户名
	Password      string // 数据库密码
	MaxIdleConns  int    `json:",default=10"`                               // 空闲中的最大连接数
	MaxOpenConns  int    `json:",default=10"`                               // 打开到数据库的最大连接数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"` // 是否开启Gorm全局日志
	LogZap        bool   // 是否通过zap写入日志文件
	SlowThreshold int64  `json:",default=1000"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + fmt.Sprintf("%d", m.Port) + ")/" + m.Dbname + "?" + m.Config
}
