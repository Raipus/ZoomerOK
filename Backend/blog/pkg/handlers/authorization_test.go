package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/handlers"
	"github.com/Raipus/ZoomerOK/blog/pkg/memory"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	r := router.SetupRouter(false)
	mockBroker := new(broker.MockBroker)
	mockMessageStore := new(memory.MockMessageStore)

	strToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzQ4NTE4MTUyLCJpZCI6MSwibG9naW4iOiJ0ZXN0dXNlciJ9.oDFERIARPY5JtbiyRomCIJJYhKfZLZ-5BQmNfDqf2zQ"
	authorizationRequest := &pb.AuthorizationRequest{
		Token: strToken,
	}
	authorizationResponse := pb.AuthorizationResponse{
		Id:             1,
		Login:          "testuser",
		Email:          "testuser@example.com",
		Token:          strToken,
		ConfirmedEmail: true,
	}
	var responseInterface interface{} = &authorizationResponse

	mockBroker.On("Authorization", authorizationRequest).Return(nil)
	mockMessageStore.On("ProcessAuthorization", authorizationRequest).Return(responseInterface, nil)

	r.Use(handlers.AuthMiddleware(mockBroker, mockMessageStore))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req, err := http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+strToken)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "success"}`, w.Body.String())
}
