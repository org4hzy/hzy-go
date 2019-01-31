package config

import (
	"log"
	"testing"
	"time"
)

// TestLFF test
func TestLFF(t *testing.T) {
	type tagData struct {
		Tm []int `json:"tm"`
		Vr []int `json:"vr"`
	}

	type tagTest struct {
		Test    string  `json:"test"`
		Version int     `json:"version"`
		Data    tagData `json:"data"`
	}

	j := &tagTest{}
	r := LoadFromFile("D:\\gowork\\hzy\\sample\\test.json", j)
	if r {
		log.Printf("%v", *j)
	} else {
		t.Error("load json from file failed ...")
	}
}

func TestCLog(t *testing.T) {
	Launcher("/d/var", "testcolorlog")

	for i := 0; i < 10; i++ {
		Info(" this is a info test .")
		Warning(" this is a warn test !")
		Error("  this is a error test !!!")
		time.Sleep(time.Second)
	}
}
