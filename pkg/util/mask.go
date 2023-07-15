package util

import (
	"regexp"
)

func Mask(str string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(str) {
		end = len(str)
	}
	mask := ""
	for i := start; i < end; i++ {
		mask += "*"
	}
	return str[:start] + mask + str[end:]
}

func MaskChineseName(name string) string {
	reg := regexp.MustCompile(`([\p{Han}]{1})([\p{Han}]*)([\p{Han}]{1})`) // 匹配中文姓名
	return reg.ReplaceAllStringFunc(name, func(s string) string {
		return Mask(s, 1, len(s)-1)
	})
}

func MaskPhone(phone string) string {
	return Mask(phone, 3, 7)
}

func MaskEmail(email string) string {
	reg := regexp.MustCompile(`^([\w\.\-]+)@([\w\-]+\.)+([\w]{2,})$`) // 匹配邮箱
	return reg.ReplaceAllStringFunc(email, func(s string) string {
		parts := reg.FindStringSubmatch(s)
		username := Mask(parts[1], 1, len(parts[1])-1)
		return username + "@" + parts[2] + parts[3]
	})
}