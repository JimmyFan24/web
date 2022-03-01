package logagent

import "web/internal/logagent/config"
//Run-->server.PrepareRun-->Run
func Run(cfg *config.Config)  error{
	server,err := createLogAgentServer(cfg)
	if err != nil{
		return err
	}
	return server.PrepareRun().Run()
}
