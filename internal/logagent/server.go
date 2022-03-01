package logagent

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"web/internal/logagent/config"
	genericoptions "web/internal/pkg/options"
	genericlogagentserver "web/internal/pkg/server"
)



type LogAgentServer struct {
	kafkaOtions *genericoptions.KafkaOptions
	gRPCLogAgentServer *gRPCLogAgentServer
	//genericAPIServer *genericoptions
	genericlogagentserver  *genericlogagentserver.GenericLogAgentServer
}
type PrepareLogAgentServer struct {
	*LogAgentServer
}

func createLogAgentServer(cfg *config.Config) (*LogAgentServer ,error){
	//开始构建log agent server
	genericConfig,err :=buildGenricConfig(cfg)

	if err != nil {
		return nil, err
	}
	extraConfig,err := buildExtraConfig(cfg)
	if err != nil{
		return nil,err
	}
	genericServer,err := genericConfig.Complete().New()
	if err != nil{
		return nil,err
	}
	extraServer,err := extraConfig.complete().New()
	if err != nil{
		return nil,err
	}

	server :=&LogAgentServer{
		genericlogagentserver: genericServer,
		gRPCLogAgentServer: extraServer,
		kafkaOtions:cfg.KafkaOptions,
	}

	return server,nil

}
type completedExtraConfig struct {
	*ExtraConfig
}
// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}
// New create a grpcAPIServer instance.
func (c *completedExtraConfig) New() (*gRPCLogAgentServer, error) {


	return &gRPCLogAgentServer{nil, c.Addr}, nil
}
// ExtraConfig defines extra configuration for the iam-apiserver.
type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	ServerCert   genericoptions.GeneratableKeyCert
	kafkaOptions *genericoptions.KafkaOptions
}
func buildExtraConfig(cfg *config.Config)(*ExtraConfig,error){
	return &ExtraConfig{
		Addr:fmt.Sprintf("%s:%d",cfg.GRPCOptions.BindAddress,cfg.GRPCOptions.BindPort),
		MaxMsgSize:   cfg.GRPCOptions.MaxMsgSize,
		ServerCert:   cfg.SecureServing.ServerCert,
		kafkaOptions: cfg.KafkaOptions,
	},nil
}
func buildGenricConfig(cfg *config.Config)(genericConfig *genericlogagentserver.Config,lastErr error){
	genericConfig = genericlogagentserver.NewConfig()

	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}



	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}



	return
}
func (l *LogAgentServer)PrepareRun()PrepareLogAgentServer{
	logrus.Info("creating PrepareLogAgentServer ")
	return PrepareLogAgentServer{l}
}
func (l *LogAgentServer)Run()error{
	logrus.Info("this is server run  func ")
	return  nil
}