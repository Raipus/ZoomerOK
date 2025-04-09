package security

import (
	"net/http"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type UserToken struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ConfirmedEmail bool   `json:"confirmed_email"`
	Image          string `json:image` // base64
}

var secretKey = []byte(config.Config.SecretKey)

func GenerateJWT(user UserToken) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["confirmed_email"] = user.ConfirmedEmail
	claims["image"] = user.Image
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Токен действителен 72 часа

	return token.SignedString(secretKey)
}

func JWTToString(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(config.Config.SecretKey))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNotSupported
		}
		return secretKey, nil
	})
}
