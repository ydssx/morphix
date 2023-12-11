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
	"time"
)

// IsPhoneNumber checks if the given string is a valid phone number.
//
// phoneNumber: the string to be checked.
// Returns: a boolean value indicating if the string is a valid phone number.
//
// Example:
//
//	IsPhoneNumber("1234567890") // true
func IsPhoneNumber(phoneNumber string) bool {
	// 定义手机号码正则表达式
	pattern := `^(1[3-9])\d{9}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phoneNumber)
}

// MD5 calculates the MD5 hash of the given text.
//
// It takes a string parameter called "text" which represents the text to be hashed.
// The function returns a string which represents the hexadecimal representation of the MD5 hash.
//
// Example:
//
//	MD5("Hello, World!") // "b10a8db164e0754105b7a99be72e3fe5"
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// IsChinese checks if the given string contains only Chinese characters.
//
// Parameter:
// str - the string to be checked.
//
// Return:
// bool - true if the string contains only Chinese characters, false otherwise.
//
// Example:
//
//	IsChinese("你好") // true
//	IsChinese("Hello") // false
func IsChinese(str string) bool {
	reg := regexp.MustCompile(`^[\u4e00-\u9fa5]+$`)
	return reg.MatchString(str)
}

// JsonToMap converts a JSON string to a map[string]interface{}.
//
// It takes a JSON string as input and returns a map[string]interface{} and an error.
//
// Example:
//
//	JsonToMap(`{"Name": "Alice", "Age": 30, "Address": "123 Main St."}`)// {"Name": "Alice", "Age": 30, "Address": "123 Main St."}
func JsonToMap(s string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	return m, err
}

// StructToMap converts a struct to a map[string]interface{}.
//
// It takes a struct as input and returns a map[string]interface{} and an error.
//
// Example:
//
//	StructToMap(struct{ Name string; Age int; Address string }) // {"Name": "Alice", "Age": 30, "Address": "123 Main St."}
func StructToMap(s interface{}) (m map[string]interface{}, err error) {
	sb, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(sb, &m)
	return
}

// Map函数，接受一个数组和一个映射函数f，返回一个新的数组
//
// Example:
//
//	Map([]int{1, 2, 3}, func(x int) int { return x + 1 }) // [2, 3, 4]
//	Map([]string{"a", "b", "c"}, func(x string) string { return strings.ToUpper(x) }) // ["A", "B", "C"]
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
//
// Example:
//
//	GenerateCode(6) // "123456"
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

// CompareRequests compares the checksum of multiple requests.
//
// It takes a slice of requests as input and returns a boolean value indicating if the checksums of all requests are equal.
//
// Example:
//
//	 type MyRequest struct {
//	  Name string
//	  Age  int
//	}
//
//	r1 := MyRequest{Name: "John", Age: 30}
//	r2 := MyRequest{Name: "John", Age: 30}
//	r3 := MyRequest{Name: "Jane", Age: 20}
//	CompareRequests(r1, r2, r3) // true
//	CompareRequests(r1, r2, r3, r3) // false
//	CompareRequests(r1, r3, r2) // true
//	CompareRequests(r1, r2, r3, r3) // false
//	CompareRequests(r1, r2, r3) // true
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

// IsZeroStruct checks if the given struct is empty.
//
// Example:
//
//	 type MyStruct struct {
//	  Name string
//	  Age  int
//	}
//
//	s := MyStruct{}
//	IsZeroStruct(s) // true
//	s.Name = "John"
//	IsZeroStruct(s) // false
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

// SetDefaults sets default values for struct fields tagged with "default"
// by reflecting over the struct. It handles setting defaults for string,
// int, float64 and bool struct fields based on the tag value.
//
// Example:
//
//	type MyStruct struct {
//	  Name string `default:"John"`
//	  Age  int    `default:"30"`
//	  Enabled bool `default:"true"`
//	}
//	SetDefaults(&MyStruct{})
//	// MyStruct will be set to:
//	MyStruct{Name: "John", Age: 30, Enabled: true}
func SetDefaults(data interface{}) {
	value := reflect.ValueOf(data).Elem()
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := typ.Field(i).Tag.Get("default")
		if tag == "" || !field.IsZero() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(tag)
		case reflect.Int:
			intValue, _ := strconv.Atoi(tag)
			field.SetInt(int64(intValue))
		case reflect.Float64:
			v, _ := strconv.ParseFloat(tag, 64)
			field.SetFloat(v)
		case reflect.Bool:
			field.SetBool(tag == "true")
		}
	}
}

func GenerateRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
