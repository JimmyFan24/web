package main

import (
	//"github.com/sirupsen/logrus"

	"web/internal/logagent"
)


func main() {
	logagent.NewLogAgent("NewLogAgent").Run()
	//logrus.Infof("test123" ,test.ABC)

}