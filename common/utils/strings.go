package utils

import (
	"regexp"
	"strings"
)

var (
	phoneExp *regexp.Regexp
)

func init() {
	// 初始化一个检查手机号的正则
	phoneExp = regexp.MustCompile(`^(1[3|4|5|8][0-9]\d{4,8})$`)

}

// 检查是否空字符串
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// 检查是否是手机号
func IsPhoneNum(phone string) bool {
	return phoneExp.MatchString(phone)
}
