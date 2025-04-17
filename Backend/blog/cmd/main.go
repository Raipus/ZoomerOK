package main

import (
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/config"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
)

var http_server = config.Config.Host + ":" + strconv.Itoa(config.Config.HttpPort)

func run_http_server() {
	router := router.SetupRouter(true)
	router.Run(http_server)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	go broker.ProductionBrokerInterface.Listen()
	go run_http_server()

	for {

	}
}
