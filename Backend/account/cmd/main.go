package main

import (
	"log"
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/broker"
	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
)

var (
	http_server = config.Config.Host + ":" + strconv.Itoa(config.Config.HttpPort)
)

func run_http_server() {
	router := router.SetupRouter(true)

	// Добавлен CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	protected := router.Group("")
	protected.Use(handlers.AuthMiddleware(postgres.ProductionPostgresInterface))
	protected.PUT(config.Config.Prefix+"/accept_friend", func(c *gin.Context) {
		handlers.AcceptFriend(c, postgres.ProductionPostgresInterface, memory.ProductionRedisInterface)
	})
	protected.POST(config.Config.Prefix+"/add_friend", func(c *gin.Context) {
		handlers.AddFriend(c, postgres.ProductionPostgresInterface)
	})
	router.PUT(config.Config.Prefix+"/change_password/:reset_link", func(c *gin.Context) {
		handlers.ChangePassword(c, postgres.ProductionPostgresInterface, caching.ProductionCachingInterface)
	})
	router.PUT(config.Config.Prefix+"/user/:login", func(c *gin.Context) {
		handlers.ChangeUser(c, postgres.ProductionPostgresInterface, memory.ProductionRedisInterface)
	})
	protected.PUT(config.Config.Prefix+"/confirm_email/:confirmation_link", func(c *gin.Context) {
		handlers.ConfirmEmail(c, postgres.ProductionPostgresInterface, caching.ProductionCachingInterface, memory.ProductionRedisInterface)
	})
	router.GET(config.Config.Prefix+"/confirm_password/:reset_link", func(c *gin.Context) {
		handlers.ConfirmPassword(c, caching.ProductionCachingInterface)
	})
	protected.DELETE(config.Config.Prefix+"/delete_friend", func(c *gin.Context) {
		handlers.DeleteFriend(c, postgres.ProductionPostgresInterface, memory.ProductionRedisInterface)
	})
	protected.DELETE(config.Config.Prefix+"/user/:login", func(c *gin.Context) {
		handlers.DeleteUser(c, postgres.ProductionPostgresInterface, memory.ProductionRedisInterface)
	})
	protected.GET(config.Config.Prefix+"/find_user", func(c *gin.Context) {
		handlers.FindUser(c, postgres.ProductionPostgresInterface, memory.ProductionRedisInterface)
	})
	protected.GET(config.Config.Prefix+"/get_friends", func(c *gin.Context) {
		handlers.GetFriends(c, memory.ProductionRedisInterface)
	})
	protected.GET(config.Config.Prefix+"/get_unaccepted_friends", func(c *gin.Context) {
		handlers.GetUnacceptedFriends(c, postgres.ProductionPostgresInterface)
	})
	protected.GET(config.Config.Prefix+"/user/:login", func(c *gin.Context) {
		handlers.GetUser(c, postgres.ProductionPostgresInterface)
	})
	router.POST(config.Config.Prefix+"/login", func(c *gin.Context) {
		handlers.Login(c, postgres.ProductionPostgresInterface, memory.ProductionRedisInterface)
	})
	router.POST(config.Config.Prefix+"/signup", func(c *gin.Context) {
		handlers.Signup(c, postgres.ProductionPostgresInterface, security.ProductionSMTPInterface, caching.ProductionCachingInterface, memory.ProductionRedisInterface)
	})
	router.POST(config.Config.Prefix+"/want_change_password", func(c *gin.Context) {
		handlers.WantChangePassword(c, postgres.ProductionPostgresInterface, security.ProductionSMTPInterface, caching.ProductionCachingInterface)
	})
	go func() {
		if err := router.Run(http_server); err != nil {
			log.Fatal("Failed to run server:", err)
		}
	}()
	log.Println("Server is running at:", http_server)
}

// TODO: linter
func main() {
	run_http_server()
	go broker.ProductionBrokerInterface.Listen()
	select {}
}
