package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
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
	return m, err
}

func StructToMap(s interface{}) (m map[string]interface{},err error) {
	sb,err:=json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(sb, &m)
	return
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
	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func CalculateChecksum(request interface{}) string {
	data := fmt.Sprintf("%v", request) // Convert request to a string
	hash := md5.Sum([]byte(data))      // Calculate MD5 hash
	return hex.EncodeToString(hash[:]) // Convert hash to a hex-encoded string
}

func CompareRequests(requests ...interface{}) bool {
	if len(requests) <= 1 {
		return true // No need to compare if there's only one request
	}

	firstChecksum := CalculateChecksum(requests[0])

	for _, request := range requests[1:] {
		checksum := CalculateChecksum(request)
		if checksum != firstChecksum {
			return false
		}
	}

	return true
}

// GenerateRandomNumber 生成指定范围内的随机整数
func GenerateRandomNumber(min, max int) int {
	if min >= max {
		panic("min must be less than max")
	}
	return rand.Intn(max-min+1) + min
}

func IsZeroStruct(s any) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if !v.Field(i).IsZero() {
				return false
			}
		}
	}
	return true
}

// SetDefaults set default value for struct field with `default` tag if the field is zero value, will panic if data is not a pointer struct.
func SetDefaults(data interface{}) {
	value := reflect.ValueOf(data).Elem()
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := typ.Field(i).Tag.Get("default")
		if tag == "" || (!field.IsZero()) {
			continue
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(tag)
		case reflect.Int:
			// You might want to handle errors here for non-integer tag values.
			// This example assumes the tags are valid integers.
			intValue, _ := strconv.Atoi(tag)
			field.SetInt(int64(intValue))
		case reflect.Float64:
			v, _ := strconv.ParseFloat(tag, 64)
			field.SetFloat(v)
		case reflect.Bool:
			if tag == "true" {
				field.SetBool(true)
			} else if tag == "false" {
				field.SetBool(false)
			}
		}
	}
}
