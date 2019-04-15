package wechat

import (
	"regexp"
	"strings"
)

func JsonDecode(users string) string {
	r, _ := regexp.Compile(`(\w+)(\s*:\s*)`)
	result := r.ReplaceAllString(users, `"$1": `)
	result = strings.ReplaceAll(result, " * 1", "")
	result = strings.ReplaceAll(result, "'", "\"")
	return result
}
