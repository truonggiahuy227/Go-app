package util

import (
	"os"

	"akawork.io/infrastructure/logger"
)

/**
 * Append Text to File
 */
func AppendTextToFile(content string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Error("Failed opening file: %s", err.Error())
		return err
	}
	defer file.Close()

	_, err = file.WriteString("\n" + content)
	if err != nil {
		logger.Error("Failed writing to file: %s", err.Error())
		return err
	}
	return nil
}

/**
 * Check File Exist
 */
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

/**
 * Returns file size
 */
func GetFileSize(filePath string) int64 {
	fi, err := os.Stat(filePath)
	if err != nil {
		logger.Error(err.Error())
	}

	return fi.Size()
}

/**
 *Cleans file
 */
func CleanFile(filePath string) {
	err := os.Truncate(filePath, 0)

	if err != nil {
		logger.Error(err.Error())
	}
}

/**
 * Deletes file
 */
func DeleteFile(filePath string) {

	var err = os.Remove(filePath)
	if err != nil {
		logger.Error(err.Error())
	}
}

/**
 * Create a new file
 */
func CreateFile(filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		logger.Fatal("[Create temp file for eventlogger]", err.Error())
	}
	f.Close()
}
