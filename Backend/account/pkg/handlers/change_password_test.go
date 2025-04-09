package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO: написать тесты для "ChangePassword"
// TODO: coverage ~ 80%
func TestChangePassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)
	r.PUT(config.Config.Prefix+"/change_password/:reset_link", func(c *gin.Context) {
		ChangePassword(c, mockPostgres, mockCache)
	})

	req, _ := http.NewRequest("PUT", config.Config.Prefix+"/change_password/:reset_link", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
