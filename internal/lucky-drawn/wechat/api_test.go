package wechat

import (
	"testing"
)

func TestGetAllData(t *testing.T) {
	GetAllData()
}

func TestFormatUsersToJson(t *testing.T) {
	filename := "./users_original.txt"
	FormatUsersToJson(filename)
}

func TestGetHeadImg(t *testing.T) {
	fakeid := "oh4N71Plyh6J8gFrN2q99aXU7_L4"
	GetHeadImg(fakeid)
}
