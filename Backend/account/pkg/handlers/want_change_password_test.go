package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO: написать тесты для "WantChangePassword"
// TODO: coverage ~ 80%
func TestWantChangePassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockSmtp := new(security.MockSmtp)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)
	r.POST(config.Config.Prefix+"/want_change_password", func(c *gin.Context) {
		WantChangePassword(c, mockPostgres, mockSmtp, mockCache)
	})

	req, _ := http.NewRequest("POST", config.Config.Prefix+"/want_change_password", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
