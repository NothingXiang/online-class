package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	phoneExp *regexp.Regexp
	emailExp *regexp.Regexp
)

func init() {
	// 初始化一个检查手机号的正则
	phoneExp = regexp.MustCompile(`^(1[3|4|5|8][0-9]\d{4,8})$`)

	// 检查邮箱的正则
	emailExp = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
}

// 检查是否空字符串
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// 检查是否是手机号
func IsPhoneNum(phone string) bool {
	return phoneExp.MatchString(phone)
}

// 检查是否邮箱地址
func IsEmailAddr(email string) bool {
	return emailExp.MatchString(email)
}

// 生成指定长度的随机验证码
func RandomCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())

	var code strings.Builder

	for length > 0 {
		fmt.Fprintf(&code, "%d", numeric[rand.Intn(len(numeric))])
		length--
	}

	return code.String()
}
