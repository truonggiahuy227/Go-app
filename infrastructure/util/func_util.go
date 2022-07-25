package util

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

/**
 * Returns a Name of interface's Func
 */
//func GetFunctionName(i interface{}) string {
//	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
//}

/**
 * Returns Name of a Func
 */
func FuncName() string {
	pc, _, line, _ := runtime.Caller(1)
	result := fmt.Sprintf("%s:%v", runtime.FuncForPC(pc).Name(), line)
	return result
}

/**
 *
 */
func BuildString(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a)
}

/**
 *
 */
func UnixToTime(value string) string {
	s, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return ""
	}
	return time.Unix(s, 0).Format(time.RFC3339Nano)
}

/**
 *
 */
func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
