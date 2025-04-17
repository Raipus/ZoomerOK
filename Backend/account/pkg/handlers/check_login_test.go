package handlers

import (
	"net/http"
)

func TestCheckLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockPostgres := new(postgres.MockPostgres)
	mockPostgres.On("CheckLogin").Return()
	router.GET("check_login", func(c *gin.Context) {
		CheckLogin(c, mockPostgres)
	})

	req, _ := http.NewRequest(http.MethodPost, "/check_login", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)
}
