package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"akawork.io/dto"
)

//Convert Nullstring to string
func ConvertNullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

/**
 * Check Md5Type
 */
func IsMd5Type(s string) bool {
	matched, _ := regexp.MatchString(`^[a-f0-9]{32}$`, s)
	return matched
}

/**
 * Converts int to string
 * @Param: {int} in
 */
func ParseString(in int) string {
	string := strconv.Itoa(in)
	return string
}

/**
 * Converts string to []string
 * @Param: {string} str
 */
func ParseArray(str string) []string {
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	res := strings.Split(str, " ")
	return res
}

/**
 * Converts []string to string
 * @Param: {[]string} str
 */
func ArraytoString(str []string) string {
	return fmt.Sprint(str)
}

func ParseMapStringInterface(istruct *dto.AccountDto) map[string]interface{} {
	var ointerface map[string]interface{}
	inrec, _ := json.Marshal(&istruct)
	json.Unmarshal(inrec, &ointerface)
	return ointerface
}
