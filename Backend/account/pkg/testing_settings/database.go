package testing_settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
)

func readJson() map[string]interface{} {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Println("data.json файл не прочтен")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func CreateData() {
	result := readJson()
	for user := range result["users"] {
		createdUser := postgres.User{
			Id:             int(user["id"]),
			Name:           user["name"],
			Email:          user["email"],
			ConfirmedEmail: user["confirmed_email"],
			Password:       user["password"],
			Phone:          user["phone"],
			City:           user["city"],
		}
		if !postgres.CreateUser(createdUser) {
			fmt.Printf("Пользователь не создан")
		}
	}
}
