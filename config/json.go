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

// ToJson 打印数据结构为json
func ToJson(d interface{}) string {


	return ""
}
