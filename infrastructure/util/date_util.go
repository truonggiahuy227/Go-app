package util

import "time"

/**
 * Returns a current date
 */
func CurrentDate() string {
	return time.Now().Format("01-02-2006")
}
