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

func TestLike(t *testing.T) {
	r := router.SetupRouter(false)

	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(1))
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	r.POST("/post/:post_id/like", func(c *gin.Context) {
		handlers.Like(c, mockPostgres)
	})

	postId := 456
	mockPostgres.On("Like", 1, postId).Return(true, nil)

	req, _ := http.NewRequest("POST", "/post/"+strconv.Itoa(postId)+"/like", nil)
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)
}
