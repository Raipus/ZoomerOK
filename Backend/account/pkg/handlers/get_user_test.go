package handlers

import (
	"net/http"
)

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockPostgres := new(postgres.MockPostgres)
	mockPostgres.On("GetUser").Return()
	router.GET("get_user", func(c *gin.Context) {
		GetUser(c, mockPostgres)
	})

	req, _ := http.NewRequest(http.MethodPost, "/get_user", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)
}
