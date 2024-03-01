package xconfig

import (
	"fmt"
)

type Redis struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func (c *Redis) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
