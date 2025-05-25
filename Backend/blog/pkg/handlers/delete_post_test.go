package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Raipus/ZoomerOK/blog/pkg/handlers"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeletePost(t *testing.T) {
	r := router.SetupRouter(false)
	userId := 1
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(userId))
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	r.DELETE("/post/:post_id", func(c *gin.Context) {
		handlers.DeletePost(c, mockPostgres)
	})

	postId := 456
	mockPostgres.On("DeletePost", userId, postId).Return(nil)

	req, _ := http.NewRequest("DELETE", "/post/"+strconv.Itoa(postId), nil)
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockPostgres.AssertExpectations(t)
}
