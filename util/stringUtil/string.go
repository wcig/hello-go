package stringUtil

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	errConvert = errors.New("failed convert")
)

func IsEmpty(s string) bool {
	return s == ""
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsBlank(s string) bool {
	return strings.Trim(s, " ") == ""
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func GetInt(s string, defaultValue int) int {
	if i, err := strToType(s, reflect.Int); err == nil && i != nil {
		if v, ok := i.(int); ok {
			return v
		}
	}
	return defaultValue
}

func GetInt32(s string, defaultValue int32) int32 {
	if i, err := strToType(s, reflect.Int32); err == nil && i != nil {
		if v, ok := i.(int32); ok {
			return v
		}
	}
	return defaultValue
}

func GetInt64(s string, defaultValue int64) int64 {
	if i, err := strToType(s, reflect.Int64); err == nil && i != nil {
		if v, ok := i.(int64); ok {
			return v
		}
	}
	return defaultValue
}

func GetFloat32(s string, defaultValue float32) float32 {
	if i, err := strToType(s, reflect.Float32); err == nil && i != nil {
		if v, ok := i.(float32); ok {
			return v
		}
	}
	return defaultValue
}

func GetFloat64(s string, defaultValue float64) float64 {
	if i, err := strToType(s, reflect.Float64); err == nil && i != nil {
		if v, ok := i.(float64); ok {
			return v
		}
	}
	return defaultValue
}

func strToType(s string, kind reflect.Kind) (interface{}, error) {
	if IsBlank(s) {
		return nil, errConvert
	}

	switch kind {
	case reflect.Int:
		return strconv.Atoi(s)
	case reflect.Int32:
		v, err := strconv.ParseInt(s, 10, 32)
		return int32(v), err
	case reflect.Int64:
		return strconv.ParseInt(s, 10, 64)
	case reflect.Float32:
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	case reflect.Float64:
		return strconv.ParseFloat(s, 64)
	default:
		return nil, errConvert
	}
}

func FormatInt(n int) string {
	return typeToStr(n, reflect.Int)
}

func FormatInt32(n int32) string {
	return typeToStr(n, reflect.Int32)
}

func FormatInt64(n int64) string {
	return typeToStr(n, reflect.Int64)
}

func FormatFloat32(n float32) string {
	return typeToStr(n, reflect.Float32)
}

func FormatFloat64(n float64) string {
	return typeToStr(n, reflect.Float64)
}

func FormatBool(n bool) string {
	return typeToStr(n, reflect.Bool)
}

func typeToStr(i interface{}, kind reflect.Kind) string {
	if i == nil {
		return ""
	}

	switch kind {
	case reflect.Int:
		return strconv.Itoa(i.(int))
	case reflect.Int32:
		return strconv.FormatInt(int64(i.(int32)), 10)
	case reflect.Int64:
		return strconv.FormatInt(i.(int64), 10)
	case reflect.Float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(i.(bool))
	default:
		return ""
	}
}
