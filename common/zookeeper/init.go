package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"strings"
	"time"
)

func New(hosts string) *zk.Conn {
	conn, _, err := zk.Connect(strings.Split(hosts, ","), 5*time.Second)
	if err != nil {
		panic(err)
	}
	return conn
}
