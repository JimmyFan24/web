package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)


const configFlagName = "config"
var cfgFile string
func init(){
	pflag.StringVarP(&cfgFile,"config","c",cfgFile,"Read configuration from specified `FILE`, "+
		"support JSON, TOML, YAML, HCL, or Java properties formats.")
}

// PrintFlags logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		logrus.Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}

func addConfigFlag(basename string,fs *pflag.FlagSet)  {
	fs.AddFlag(pflag.Lookup(configFlagName))

	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename), "-", "_", -1))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cobra.OnInitialize(func() {
		logrus.Info("this is cobra.Oninit func ")
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)

		}else {
			//如果没有配置-c，没有设定配置文件，那么使用默认的配置文件
			viper.AddConfigPath(".")

			viper.SetConfigName("config")
		}
		if err:= viper.ReadInConfig();err!=nil{
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
			os.Exit(1)
		}
		logrus.Infof("--testing read viper config:%s",viper.GetString("abc"))
		logrus.Infof("----Used configuration file is: %s\n", viper.ConfigFileUsed())
	})

}
