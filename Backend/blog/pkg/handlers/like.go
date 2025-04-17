package handlers

import (
	"github.com/Raipus/ZoomerOK/account/pkg/broker"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type LikeForm struct {
	PostId int
}

func Like(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	postId := c.Param("post_id")

}
