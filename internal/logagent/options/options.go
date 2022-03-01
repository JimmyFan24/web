package options

import (
	"github.com/sirupsen/logrus"

	//"github.com/spf13/pflag"
	genericoptions "web/internal/pkg/options"
	"web/pkg/app"
)

type Options struct {
	MysqlOptions *genericoptions.MySQLOptions
	KafkaOptions *genericoptions.KafkaOptions
	EtcdOptions *genericoptions.EtcdOptions
	GenericServerRunOptions *genericoptions.ServerRunOptions
	SecureServing *genericoptions.SecureServingOptions
	GRPCOptions *genericoptions.GRPCOptions
}

func (o *Options) Validate() []error {
	return nil
}

func NewOptions() *Options {
	logrus.Info("2.创建新的Options，这里包括mysql和kafka,etcd")
	return &Options{
		MysqlOptions: genericoptions.NewMysqlOptions(),
		KafkaOptions: genericoptions.NewKafkaOptions(),
		EtcdOptions: genericoptions.NewEtcdOptions(),
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		SecureServing: genericoptions.NewSecureServingOptions(),
		GRPCOptions: genericoptions.NewGRPCOptions(),
	}
}

func (o *Options)Flags() (fs app.NamedFlagSet) {

	o.MysqlOptions.AddFlags(fs.FlagSet("mysql"))
	o.KafkaOptions.AddFlags(fs.FlagSet("kafka"))
	o.EtcdOptions.AddFlags(fs.FlagSet("etcd"))
	o.GenericServerRunOptions.AddFlags(fs.FlagSet("generic"))
	o.SecureServing.AddFlags(fs.FlagSet("secure serving"))
	o.GRPCOptions.AddFlags(fs.FlagSet("grpc"))
	//logrus.Infof("初始化flagset完成,kafka :%v,etcd:%v",fs.FlagSet("kafka"),fs.FlagSet("etcd"))

	return fs
}
