package main

import (
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
)

var (
	http_server = config.Config.Host + ":" + strconv.Itoa(config.Config.HttpPort)
)

func run_http_server() {
	router := router.SetupRouter(true)
	router.POST(config.Config.Prefix+"/signup", handlers.Signup)
	router.POST(config.Config.Prefix+"/login", handlers.Login)
	router.GET(config.Config.Prefix+"/confirm_email/:url", handlers.ConfirmEmail)
	router.GET(config.Config.Prefix+"/confirm_password/:url", handlers.ConfirmPassword)
	router.Run(http_server)
}

// TODO: linter
func main() {
	postgres.Init()
	postgres.Migrate()

	run_http_server()
}
