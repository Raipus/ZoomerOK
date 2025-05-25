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

func TestDeleteComment(t *testing.T) {
	r := router.SetupRouter(false)
	userId := 1
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(1))
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	r.DELETE("/post/:post_id/comments/:comment_id", func(c *gin.Context) {
		handlers.DeleteComment(c, mockPostgres)
	})

	commentId := 123
	mockPostgres.On("DeleteComment", userId, commentId).Return(nil)

	req, _ := http.NewRequest("DELETE", "/post/1/comments/"+strconv.Itoa(commentId), nil)
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockPostgres.AssertExpectations(t)
}
