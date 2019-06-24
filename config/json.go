package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// LoadFromFile 文件加载json
func LoadFromFile(fp string, j interface{}) bool {
	bytes, e := ioutil.ReadFile(fp)
	if e != nil {
		log.Printf("LoadFromFile | %v", e)
		return false
	}

	e = json.Unmarshal(bytes, j)
	if e != nil {
		log.Printf("LoadFromFile %v", e.Error())
		return false
	}

	return true
}

// ToJSON only accesses the exported fileds of struct types(an uppercase letter.)
func ToJSON(d interface{}) string {
	bytes, err := json.Marshal(d)
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}
