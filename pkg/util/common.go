package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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
//	IsPhoneNumber("1234567890") // false
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

// CalculateChecksum 计算给定请求的校验和
// 它将请求转换为字符串,计算 MD5 哈希,并将哈希转换为十六进制编码的字符串
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
		default:
			panic(fmt.Sprintf("unsupported type: %s", field.Kind()))
		}
	}
}

// GenerateRandomString generates a random string of the given length.
// It does this by selecting random bytes from the set of alphanumeric characters.
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

// InterfaceToString将任意类型的数据转换为字符串
// 根据数据类型,使用不同的方式进行转换
//   - string类型直接返回
//   - 整型使用strconv.Itoa转换为字符串
//   - 浮点数使用strconv.FormatFloat转换为字符串
//   - 其他类型使用fmt.Sprintf转换为字符串
func InterfaceToString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case int:
		return strconv.Itoa(data.(int))
	case int8:
		return strconv.Itoa(int(data.(int8)))
	case int16:
		return strconv.Itoa(int(data.(int16)))
	case int32:
		return strconv.Itoa(int(data.(int32)))
	case int64:
		return strconv.Itoa(int(data.(int64)))
	case uint:
		return strconv.Itoa(int(data.(uint)))
	case uint8:
		return strconv.Itoa(int(data.(uint8)))
	case uint16:
		return strconv.Itoa(int(data.(uint16)))
	case uint32:
		return strconv.Itoa(int(data.(uint32)))
	case uint64:
		return strconv.Itoa(int(data.(uint64)))
	case float32:
		return strconv.FormatFloat(float64(data.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(data.(float64), 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", data)
	}
}

// InterfaceToInt将任意类型的数据转换为int类型
// 根据数据类型,使用不同的方式进行转换:
//   - string类型使用strconv.Atoi转换为int
//   - 整型和浮点型使用类型转换转换为int
//   - 其他类型返回0
func InterfaceToInt(data interface{}) int {
	switch data.(type) {
	case string:
		v, _ := strconv.Atoi(data.(string))
		return v
	case int:
		return data.(int)
	case int8:
		return int(data.(int8))
	case int16:
		return int(data.(int16))
	case int32:
		return int(data.(int32))
	case int64:
		return int(data.(int64))
	case uint:
		return int(data.(uint))
	case uint8:
		return int(data.(uint8))
	case uint16:
		return int(data.(uint16))
	case uint32:
		return int(data.(uint32))
	case uint64:
		return int(data.(uint64))
	case float32:
		return int(data.(float32))
	case float64:
		return int(data.(float64))
	default:
		return 0
	}
}

// InterfaceToFloat64 converts the given data of type interface{} to a float64.
//
// The data can be of type string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32,
// uint64, float32, or float64. If the data is of type string, it is parsed to a float64 using
// the strconv.ParseFloat function. For other numeric types, they are directly converted to
// float64. If the data is not one of the supported types, the function returns 0.
//
// Parameters:
//   - data: The data to be converted to float64.
//
// Returns:
//   - float64: The converted value of data as a float64.
func InterfaceToFloat64(data interface{}) float64 {
	switch data.(type) {
	case string:
		v, _ := strconv.ParseFloat(data.(string), 64)
		return v
	case int:
		return float64(data.(int))
	case int8:
		return float64(data.(int8))
	case int16:
		return float64(data.(int16))
	case int32:
		return float64(data.(int32))
	case int64:
		return float64(data.(int64))
	case uint:
		return float64(data.(uint))
	case uint8:
		return float64(data.(uint8))
	case uint16:
		return float64(data.(uint16))
	case uint32:
		return float64(data.(uint32))
	case uint64:
		return float64(data.(uint64))
	case float32:
		return float64(data.(float32))
	case float64:
		return data.(float64)
	default:
		return 0
	}
}

// GetEnv returns the value of the environment variable named by the key.
// If the environment variable is not present, the fallback value is returned instead.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

