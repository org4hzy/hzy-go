package test

import (
	"testing"

	"github.com/hzy/util"
)

//TestBase64 test base64
func TestBase64(t *testing.T) {
	r := util.Base64("1234567890")
	t.Error(r)
}
