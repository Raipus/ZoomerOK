package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	loginData := LoginForm{
		LoginOrEmail: "testuser",
		Password:     "password",
	}

	r.POST("/login", func(c *gin.Context) {
		Login(c, mockPostgres)
	})

	var token string = "new-token"
	mockPostgres.On("Login", loginData.LoginOrEmail, loginData.Password).Return(token, "")

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := gin.H{
		"token": token,
	}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
	mockPostgres.AssertExpectations(t)
}
