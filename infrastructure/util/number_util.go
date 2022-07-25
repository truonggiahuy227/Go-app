package util

import (
	"strconv"

	"akawork.io/infrastructure/logger"
)

/**
 * Converts string to int
 * @Param: {string} str
 */
func ParseInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		logger.Error("[ParserInt] - Error: %s, str: %s", err, str)
		return -1
	}
	return number
}

/**
 * Converts string to int64
 * @Param: {string} str
 */
func ParseInt64(str string) int64 {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return -1
	}
	return number
}

/**
 *
 */
func ToString(i int) string {
	return strconv.Itoa(i)
}

/**
 *
 */
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

/**
 *
 */
func FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 6, 64)
}

/**
 *
 */
func InttoUint32(n int) uint32 {
	return uint32(n)
}
