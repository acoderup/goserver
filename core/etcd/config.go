package etcd

import (
	"fmt"
	"time"

	"github.com/acoderup/goserver/core"
)

var Config = Configuration{}

type Configuration struct {
	Url         []string
	UserName    string
	Password    string
	DialTimeout int // second
}

func (c *Configuration) Name() string {
	return "etcd"
}

func (c *Configuration) Init() error {
	err := globalClient.Open(c.Url, c.UserName, c.Password, time.Duration(c.DialTimeout)*time.Second)
	if err != nil {
		panic(fmt.Sprintf("etcd init error:%v", err))
	}
	return nil
}

func (c *Configuration) Close() error {
	if globalClient != nil {
		globalClient.Close()
	}
	return nil
}

func init() {
	core.RegistePackage(&Config)
}
