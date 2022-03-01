package options

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gorm.io/gorm"
	"time"
	"web/pkg/db"
)

type MySQLOptions struct {
	Host     string `json:"host,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
	LogLevel int    `json:"log_level,omitempty"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"   "`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"   "`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" "`

}
func NewMysqlOptions() *MySQLOptions{
	logrus.Info("3.创建新的Options，这里是mysql")
	return &MySQLOptions{
		Host:     "127.0.0.1:3306",
		Username: "",
		Password: "",
		Database: "",
		LogLevel: 1,
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}
// Validate verifies flags passed to MySQLOptions.
func (o *MySQLOptions) Validate() []error {
	errs := []error{}

	return errs
}
func(o *MySQLOptions)AddFlags(fs *pflag.FlagSet){
	fs.StringVar(&o.Host,"mysql.host",o.Host,""+
		"MySQL service host address. If left blank, the following related mysql options will be ignored.")
	fs.StringVar(&o.Username,"mysql.username",o.Username,""+
		"Username for access to mysql service.")
	fs.StringVar(&o.Password, "mysql.password", o.Password, ""+
		"Password for access to mysql, should be used pair with password.")

	fs.StringVar(&o.Database, "mysql.database", o.Database, ""+
		"Database name for the server to use.")
	fs.IntVar(&o.LogLevel, "mysql.log-mode", o.LogLevel, ""+
		"Specify gorm log level.")

	fs.IntVar(&o.MaxIdleConnections, "mysql.max-idle-connections", o.MaxOpenConnections, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&o.MaxOpenConnections, "mysql.max-open-connections", o.MaxOpenConnections, ""+
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&o.MaxConnectionLifeTime, "mysql.max-connection-life-time", o.MaxConnectionLifeTime, ""+
		"Maximum connection life time allowed to connecto to mysql.")
}

func (o *MySQLOptions)NewClient()(*gorm.DB,error)  {
	opts :=&db.Options{
		Host:                  o.Host,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		LogLevel:              o.LogLevel,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,

	}

	return db.NewMysqlCli(opts)
}