package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Translate(key string) string {
	m := make(map[string]map[string]interface{})
	lang := os.Getenv("lang")

	jsonFile, err := os.Open("./src/assets/" + lang + ".json")
	if err != nil {
		jsonFile, _ = os.Open("./src/assets/en.json")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &m)
	if err != nil {
		fmt.Println("can`t Unmarshal file")
	}

	expression := strings.Split(key, ".")
	if len(expression) != 2 {
		return key
	}

	value, ok := m[expression[0]]
	if ok {
		return value[expression[1]].(string)
	} else {
		return key
	}
}
