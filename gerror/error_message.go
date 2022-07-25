package gerror

/**
 * Translate
 */
func T(errorCode uint32) string {
	switch errorCode {
	//////////////////////////
	// Client-side
	//////////////////////////
	case ErrorBindData:
		return "Failed to bind data"
	case ErrorValidData:
		return "Failed to valid data"
	case ErrorRedis:
		return "Error when use Redis"
		//////////////////////////
	// Server-side
	//////////////////////////
	case ErrorConnect:
		return "Failed to connect database"
	case ErrorSaveData:
		return "Failed to save data"
	case ErrorRetrieveData:
		return "Failed to retrieve data"
	}

	return "Unknown error"
}
