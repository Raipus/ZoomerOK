package main

import (
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
)

var (
	http_server = config.Config.Host + ":" + strconv.Itoa(config.Config.HttpPort)
)

func run_http_server() {
	router := router.SetupRouter(true)
	router.POST(config.Config.Prefix+"/signup", func(c *gin.Context) {
		handlers.Signup(c, postgres.ProductionPostgresInterface)
	})
	router.POST(config.Config.Prefix+"/login", func(c *gin.Context) {
		handlers.Login(c, postgres.ProductionPostgresInterface)
	})
	// router.PUT(config.Config.Prefix+"/change_password", handlers.ChangePassword)
	router.GET(config.Config.Prefix+"/confirm_email/:confirmation_link", func(c *gin.Context) {
		handlers.ConfirmEmail(c, caching.ProductionCachingInterface)
	})
	router.GET(config.Config.Prefix+"/confirm_password/:reset_link", func(c *gin.Context) {
		handlers.ConfirmPassword(c, caching.ProductionCachingInterface)
	})
	router.Run(http_server)
}

// TODO: linter
func main() {
	run_http_server()
}
