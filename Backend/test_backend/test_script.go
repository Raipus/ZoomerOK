package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SignupForm struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

const AccountUrl string = "http://localhost:8080/"
const BlogUrl string = "http://localhost:1337/blog/"

func main() {
	signupData := SignupForm{
		Login:    "testuser",
		Name:     "Тестовый пользователь",
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonData, err := json.Marshal(signupData)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	resp, err := http.Post(AccountUrl+"signup", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		var response SignupResponse
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Println("Ошибка при декодировании ответа:", err)
			return
		}
		fmt.Println("Регистрация прошла успешно! Токен:", response.Token)
	} else {
		var response SignupResponse
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Println("Ошибка при декодировании ответа:", err)
			return
		}
		fmt.Printf("Ошибка регистрации: %s\n", resp.Status)
	}
}
