package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func IsPhoneNumber(phoneNumber string) bool {
	// 定义手机号码正则表达式
	pattern := `^(1[3-9])\d{9}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phoneNumber)
}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func IsChinese(str string) bool {
	reg := regexp.MustCompile(`^[\u4e00-\u9fa5]+$`)
	return reg.MatchString(str)
}

func JsonToMap(s string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Map函数，接受一个数组和一个映射函数f，返回一个新的数组
func Map[T any](nums []T, f func(T) T) []T {
	result := make([]T, len(nums))
	for i, num := range nums {
		result[i] = f(num)
	}
	return result
}

// Reduce函数，接受一个整数数组和一个归约函数f，返回归约结果
func Reduce(nums []int, f func(int, int) int, init int) int {
	result := init
	for _, num := range nums {
		result = f(result, num)
	}
	return result
}

// 生成指定长度的随机数字字符串
func GenerateCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}
