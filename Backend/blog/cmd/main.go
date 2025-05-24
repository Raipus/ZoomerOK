package main

import (
	"log"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/config"
	"github.com/Raipus/ZoomerOK/blog/pkg/handlers"
	"github.com/Raipus/ZoomerOK/blog/pkg/memory"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
)

var http_server = config.Config.Host + ":" + strconv.Itoa(config.Config.HttpPort)

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
	protected.Use(handlers.AuthMiddleware(broker.ProductionBrokerInterface, memory.ProductionMessageStore))
	protected.POST(config.Config.Prefix+"/post/:post_id/create_comment", func(c *gin.Context) {
		handlers.CreateComment(c, postgres.ProductionPostgresInterface)
	})
	protected.POST(config.Config.Prefix+"/create_post", func(c *gin.Context) {
		handlers.CreatePost(c, postgres.ProductionPostgresInterface)
	})
	protected.GET(config.Config.Prefix+"/post/:post_id", func(c *gin.Context) {
		handlers.GetPost(c, postgres.ProductionPostgresInterface, broker.ProductionBrokerInterface, memory.ProductionMessageStore)
	})
	protected.DELETE(config.Config.Prefix+"/post/:post_id/comments/:comment_id", func(c *gin.Context) {
		handlers.DeleteComment(c, postgres.ProductionPostgresInterface)
	})
	protected.DELETE(config.Config.Prefix+"/post/:post_id", func(c *gin.Context) {
		handlers.DeletePost(c, postgres.ProductionPostgresInterface)
	})
	protected.GET(config.Config.Prefix+"/post/:post_id/comments", func(c *gin.Context) {
		handlers.GetComments(c, postgres.ProductionPostgresInterface, broker.ProductionBrokerInterface, memory.ProductionMessageStore)
	})
	protected.GET(config.Config.Prefix+"/posts", func(c *gin.Context) {
		handlers.GetPosts(c, postgres.ProductionPostgresInterface, broker.ProductionBrokerInterface, memory.ProductionMessageStore)
	})
	protected.GET(config.Config.Prefix+"/user/:id/posts", func(c *gin.Context) {
		handlers.GetUserPosts(c, postgres.ProductionPostgresInterface)
	})
	protected.POST(config.Config.Prefix+"/post/:post_id/like", func(c *gin.Context) {
		handlers.Like(c, postgres.ProductionPostgresInterface)
	})
	go func() {
		if err := router.Run(http_server); err != nil {
			log.Fatal("Failed to run server:", err)
		}
	}()
	log.Println("Server is running at:", http_server)
}

func main() {
	run_http_server()
	go broker.ProductionBrokerInterface.Listen()
	select {}
}
