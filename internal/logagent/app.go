package logagent

import (
	"github.com/sirupsen/logrus"
	"web/internal/logagent/config"
	"web/internal/logagent/options"
	"web/pkg/app"
)

func init()  {
	//logrus.SetLevel(logrus.WarnLevel)
}
func NewLogAgent(basename string)  *app.App{
	logrus.Info("1.开始创建新的logagent实例...")
	opts := options.NewOptions()
	logrus.Info("5.创建完mysql和kafka的options实例，开始")
	application := app.NewApp("logAgentApp----",basename,
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),)
	logrus.Infof("this is info of app item%v",application)
	return application
}
func run(opt *options.Options) app.RunFunc{
	return func(basename string) error {

		logrus.Info("logagent server is running...")
		//fmt.Printf("testing read viper config:%v\n",viper.GetString("KafkaOptions.Host"))
		//.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
		//构建应用运行的配置信息
		cfg,err := config.CreateConfigFromOptions(opt)
		if err != nil{
			return err
		}
		return Run(cfg)
	}

}