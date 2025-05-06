package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
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
	mockRedis := new(memory.MockRedis)

	signupData := handlers.SignupForm{
		Login:    "testuser",
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "securepassword",
	}

	birthday := time.Now()
	byteImage := config.Config.Photo.ByteImage
	var token string = "new-token"
	user := postgres.User{
		Id:             1,
		Login:          "testuser",
		Name:           "Тестовый Пользователь",
		Email:          "testuser@example.com",
		ConfirmedEmail: true,
		Password:       "securepassword",
		Birthday:       &birthday,
		Phone:          "123-456-7890",
		City:           "Москва",
		Image:          byteImage,
	}
	redisAuthorization := memory.RedisAuthorization{
		UserId: user.Id,
		Token:  token,
		Login:  user.Login,
		Email:  user.Email,
	}
	redisUser := memory.RedisUser{
		UserId: user.Id,
		Login:  user.Login,
		Name:   user.Name,
		Image:  config.Config.Photo.Base64Small,
	}

	mockSmtp.On("SendConfirmEmail", signupData.Login, signupData.Email, mockCache).Return(nil)
	mockPostgres.On("GetUserByLogin", signupData.Login).Return(postgres.User{})
	mockPostgres.On("GetUserByEmail", signupData.Email).Return(postgres.User{})
	mockPostgres.On("Signup", signupData.Login, signupData.Name, signupData.Email, signupData.Password).Return(user, token, true)
	mockRedis.On("SetAuthorization", redisAuthorization).Return()
	mockRedis.On("SetUser", redisUser).Return()

	jsonData, err := json.Marshal(signupData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.POST("/signup", func(c *gin.Context) {
		handlers.Signup(c, mockPostgres, mockSmtp, mockCache, mockRedis)
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
	mockRedis := new(memory.MockRedis)

	signupData := handlers.SignupForm{
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
		handlers.Signup(c, mockPostgres, mockSmtp, mockCache, mockRedis)
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
	mockRedis.AssertExpectations(t)
}

func TestSignupEmailExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockSmtp := new(security.MockSmtp)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)
	mockRedis := new(memory.MockRedis)

	signupData := handlers.SignupForm{
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
		handlers.Signup(c, mockPostgres, mockSmtp, mockCache, mockRedis)
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
	mockRedis.AssertExpectations(t)
}
