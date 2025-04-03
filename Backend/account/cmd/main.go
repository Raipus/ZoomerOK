package main

import (
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

var (
	http_server = config.Config.Host + ":" + strconv.Itoa(config.Config.HttpPort)
)

func run_http_server() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	// router.POST(config.Config.Prefix+"/registry", handlers.Registry)
	router.POST(config.Config.Prefix+"/login", handlers.Login)
	router.Run(http_server)
}

func main() {
	postgres.Init()
	postgres.Migrate()

	run_http_server()
}
