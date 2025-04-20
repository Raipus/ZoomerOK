package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockSmtp := new(security.MockSmtp)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	signupData := SignupForm{
		Login:    "testuser",
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "securepassword",
	}

	var token string = "new-token"
	mockSmtp.On("SendConfirmEmail", signupData.Name, signupData.Email, mockCache).Return(nil)
	mockPostgres.On("GetUserByLogin", signupData.Login).Return(postgres.User{})
	mockPostgres.On("GetUserByEmail", signupData.Email).Return(postgres.User{})
	mockPostgres.On("Signup", signupData.Login, signupData.Name, signupData.Email, signupData.Password).Return(token, true)

	jsonData, err := json.Marshal(signupData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.POST("/signup", func(c *gin.Context) {
		Signup(c, mockPostgres, mockSmtp, mockCache)
	})

	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonData))
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
	mockSmtp.AssertExpectations(t)
}

func TestSignupLoginExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockSmtp := new(security.MockSmtp)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	signupData := SignupForm{
		Login:    "existinguser",
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "securepassword",
	}

	// Настройка моков для существующего пользователя
	mockPostgres.On("GetUserByLogin", signupData.Login).Return(postgres.User{Login: signupData.Login})

	jsonData, err := json.Marshal(signupData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.POST("/signup", func(c *gin.Context) {
		Signup(c, mockPostgres, mockSmtp, mockCache)
	})

	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	expectedResponse := gin.H{"error": "Логин 'existinguser' уже существует"}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	mockPostgres.AssertExpectations(t)
}

func TestSignupEmailExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockSmtp := new(security.MockSmtp)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	signupData := SignupForm{
		Login:    "newuser",
		Name:     "Test User",
		Email:    "existing@example.com",
		Password: "securepassword",
	}

	// Настройка моков для существующего email
	mockPostgres.On("GetUserByLogin", signupData.Login).Return(postgres.User{})
	mockPostgres.On("GetUserByEmail", signupData.Email).Return(postgres.User{Email: signupData.Email})

	jsonData, err := json.Marshal(signupData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.POST("/signup", func(c *gin.Context) {
		Signup(c, mockPostgres, mockSmtp, mockCache)
	})

	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	expectedResponse := gin.H{"error": "Электронная почта 'existing@example.com' уже существует"}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	mockPostgres.AssertExpectations(t)
}
