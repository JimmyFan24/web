package db

import (
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
)


//初始化并且连接etcd
// New create a new gorm db instance with the given options.
func NewEtcdCli (address []string)(cli *clientv3.Client,err error){
	cli,err = clientv3.New(clientv3.Config{
		Endpoints: address,
		DialTimeout: time.Second*5,
	})

	if err != nil{
		return nil,err
	}
	logrus.Info("etcd 初始化成功...")
	return cli,nil
}
