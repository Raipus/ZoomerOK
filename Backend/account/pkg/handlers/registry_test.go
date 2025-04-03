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
func TestRegistry(t *testing.T) {
	r := router.SetupRouter(false)
	r.POST("/"+config.Config.Prefix+"/registry", Registry)

	req, _ := http.NewRequest("POST", config.Config.Prefix+"/registry", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
