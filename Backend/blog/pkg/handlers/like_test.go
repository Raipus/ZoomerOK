package handlers

import (
	"new/http"
)

func TestLike(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	r.POST("/like")
}
