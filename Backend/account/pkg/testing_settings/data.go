package testing_settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadJson() map[string]interface{} {
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
