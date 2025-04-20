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

func TestDeleteComment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	r.DELETE("/comment/:comment_id", func(c *gin.Context) {
		DeleteComment(c, mockPostgres)
	})

	var commentId int = 1
	mockPostgres.On("DeleteComment", 1, commentId).Return(nil)

	req, err := http.NewRequest("DELETE", "/comment/"+strconv.Itoa(commentId), nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	mockPostgres.AssertExpectations(t)
}
