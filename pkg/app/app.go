package app

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	commands    []*Command
	args        cobra.PositionalArgs
	cmd         *cobra.Command

}
type RunFunc func(basename string)error
type Options func(a *App)

func WithRunFunc(run RunFunc) Options  {
	return func(a *App) {
		a.runFunc = run
	}
}
func WithNoConfig() Options  {
	return func(a *App) {
		a.noConfig = true
	}
}
func WithOptions(opt CliOptions) Options {
	return func(a *App) {
		a.options = opt
		//fmt.Println()
	}
}
func NewApp(name string ,basename string,opts ...Options) *App {
	logrus.Info("准备新建app实例...")
	a := &App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(a)
	}
	a.buildCommand()
	return a
}
func(a *App)buildCommand(){
	logrus.Info("开始构建应用命令行")
	cmd:= &cobra.Command{
		Use:a.basename,
		Short: a.name,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args: a.args,
	}
	// cmd.SetUsageTemplate(usageTemplate)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	//fmt.Println("this is cmd flags",cmd.Flags().SortFlags)
	if len(a.commands) > 0{
		//fmt.Println("this is buildcommand func--",a.commands)
		for _,command :=range a.commands{
			cmd.AddCommand(command.cobraCommand())
		}
	}
	//fmt.Println("this is buildcommand func",a.commands)
	if a.runFunc != nil{
		logrus.Info("this is a.runcommand..")
		cmd.RunE = a.runCommand
	}

	var namedFlagSets NamedFlagSet
	if a.options != nil {
		//这里返回的是带命名的flagset集合，比如说mysql和kafka的flagset
		logrus.Info("初始化kafka和mysql的flagset")
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		//其实这里也就两个flagset，一个kafka一个mysql
		//fmt.Println("this is flagset len ",len(namedFlagSets.FlagSets))
		for _,f := range namedFlagSets.FlagSets{
			fs.AddFlagSet(f)
		}

	}
	if !a.noConfig{
		logrus.Infof("this is into addconfig func ")
		addConfigFlag(a.basename,namedFlagSets.FlagSet("global"))
	}
	a.cmd = cmd

}

func (a *App)runCommand(cmd *cobra.Command,args []string)error{
	// run application
	//PrintFlags(cmd.Flags())
	logrus.Info("this is runcommand func running")

	if !a.noConfig {
		logrus.Infof("kafka server is :%v",viper.Get("Host"))
		if err:= viper.BindPFlags(cmd.Flags());err!=nil{
			logrus.Errorf("viper bind pflags failed.%v",err)

			return err
		}

		if err := viper.Unmarshal(a.options);err!= nil{
			logrus.Errorf("viper Unmarshal failed.%v",err)
			return err
		}

	}
	//PrintFlags(a.options.Flags().FlagSets["mysql"])
	//logrus.Infof("get mysql config from bindflags:%s\n",viper.Get("mysql.database"))
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}
	return nil
}
// Run is used to launch the application.
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}