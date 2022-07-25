package util

import (
	"encoding/json"

	"akawork.io/infrastructure/logger"
)

/**
 * Converts an object to string
 * @Param {object} v
 * @Return {string}
 */
func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		// fmt.Println(err)
		return ""
	}

	return string(b)
}

/**
 *
 */
func JsonMarshal(v interface{}) ([]byte, error) {
	byte, err := json.Marshal(v)
	if err != nil {
		logger.Error("Marshal Exception: %s, Data: %v", err, v)
	}
	return byte, err
}
