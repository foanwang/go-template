package main

import (
	"os"
	"testing/src/config"
	"testing/src/router"
	"testing/src/util"
)

func main() {
	dir, _ := os.Getwd()
	util.InitProperty(dir + "/src/resource/config.yml")
	r := router.SetRouter() //gin
	config.InitLog()        //log to elk
	config.InitDB()         //mysql
	config.InitRedis()      //redis
	// userRepo := repository.NewUserRepository(config.GetDB())

	r.Run(":" + util.GetElement("port"))
}
