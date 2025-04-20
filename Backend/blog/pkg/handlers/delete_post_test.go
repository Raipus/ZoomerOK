package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	r.DELETE("/post/:post_id", func(c *gin.Context) {
		DeletePost(c, mockPostgres)
	})

	var postId int = 1
	mockPostgres.On("DeletePost", 1, postId).Return(nil)

	req, err := http.NewRequest("DELETE", "/post/"+strconv.Itoa(postId), nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	mockPostgres.AssertExpectations(t)
}
