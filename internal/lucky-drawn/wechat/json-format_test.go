package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	b,err := ioutil.ReadFile("./users_original.txt")
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(JsonDecode(string(b)))
	users := &Users{}
	json.Unmarshal([]byte(JsonDecode(string(b))),users)
	fmt.Println(fmt.Sprintf("Users %+v",users))
}
