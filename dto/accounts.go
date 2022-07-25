package dto

type AccountDto struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type RequestsDto struct {
	TimeInSecond string `json:"TimeInSecond"`
	Count        int    `json:"Count"`
}

type ColorDto struct {
	Flag string `json:"Flag"`
}

type HealthCheckDto struct {
	Status string `json:"Status"`
}
