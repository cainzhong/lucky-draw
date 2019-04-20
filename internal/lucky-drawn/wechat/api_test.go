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
