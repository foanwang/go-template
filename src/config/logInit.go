package config

import (
	"fmt"
	"net"
	util "testing/src/util"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func GetLog() *logrus.Logger {
	if log == nil {
		InitLog()
	}
	return log
}

func InitLog() {
	log := logrus.New()
	conn, err := net.Dial("tcp", util.GetElement("logstash.host")+":"+util.GetElement("logstash.port"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("log init......")
	hook := logrustash.New(conn, logrustash.DefaultFormatter(
		logrus.Fields{
			"service": util.GetElement("service.name")}))
	log.Hooks.Add(hook)

}
