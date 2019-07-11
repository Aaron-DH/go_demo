package main

import (
	"errors"
	"net/http"
	"time"

	"demo1_gogin_api/config"
	"demo1_gogin_api/log"
	"demo1_gogin_api/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

// @title Apiserver Example API
// @version 1.0
// @description apiserver demo

// @contact.name Aaron
// @contact.url http://www.swagger.io/support
// @contact.email 344677472@qq.com

// @host localhost:8080/swagger/index.html
// @BasePath /v1
func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,

		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.RunLog.Critical("The router has no response, or it might took too long to start up.", err)
		}
		log.RunLog.Info("The API Server has been deployed successfully.")
	}()

	log.RunLog.Info("Start to listen httpserver on: %s", viper.GetString("addr"))
	log.RunLog.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("checkurl"))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.RunLog.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
