package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var setting map[interface{}]interface{}

func InitProperty(path string) {
	fmt.Println("init setting...")
	data, err := ioutil.ReadFile(path)
	// m := make(map[interface{}]interface{})
	if err != nil {
		fmt.Println("read file error")
	}
	err = yaml.Unmarshal([]byte(data), &setting)
	if err != nil {
		fmt.Println("yaml paser error")
	}

}

func GetElement(key string) string {
	if value, ok := setting[key]; ok {
		return fmt.Sprint(value)
	}
	return ""
}
