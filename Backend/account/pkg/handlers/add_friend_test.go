package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddFriend(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockPostgres := new(postgres.MockPostgres)
	mockPostgres.On("AcceptFriendRequest").Return()
	router.POST("/add_friend", func(c *gin.Context) {
		AddFriend(c, mockPostgres)
	})

	req, _ := http.NewRequest(http.MethodPost, "/add_friend", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)
}
