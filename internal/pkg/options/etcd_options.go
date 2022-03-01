package options

import (
	"github.com/spf13/pflag"
	"go.etcd.io/etcd/clientv3"
)

type EtcdOptions struct {
	Address string `json:"address"`
	EtcdClient *clientv3.Client
}

func NewEtcdOptions() *EtcdOptions{
	return &EtcdOptions{Address: ""}
}

func (e *EtcdOptions)AddFlags(fs *pflag.FlagSet)  {
	fs.StringVarP(&e.Address,"etcd-server","s","","etcd server ip:port")
}
/*
func (e *EtcdOptions)NewEtcdClient(opt *EtcdOptions)error{
	addr := opt.Address
	cli ,err := db.NewEtcdCli([]string{addr})

	if err != nil{
		return err
	}
	e.EtcdClient = cli
	return nil
}*/