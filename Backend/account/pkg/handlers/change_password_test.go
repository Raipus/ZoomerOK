package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/stretchr/testify/assert"
)

// TODO: написать тесты для "registry"
// TODO: coverage ~ 80%
// TODO: test database?
// TODO: MOCK STMP
func TestSignup(t *testing.T) {
	r := router.SetupRouter(false)
	r.POST("/"+config.Config.Prefix+"/change_password", ChangePassword)

	req, _ := http.NewRequest("PUT", config.Config.Prefix+"/change_password", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
