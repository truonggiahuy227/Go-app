package main

import (
	"fmt"
	"os"
	"time"

	"akawork.io/infrastructure/cache"
	"akawork.io/infrastructure/logger"
	"akawork.io/module/account"
	"akawork.io/module/homepage"
	"akawork.io/module/initialize"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`Debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	} else {
		fmt.Println("Service RUN on PRODUCTION mode")
	}
}

func main() {
	/********************************************************************/
	/* CONFIGURE LOG                          */
	/********************************************************************/
	logPath := viper.GetString("Log.Path")
	logPrefix := viper.GetString("Log.Prefix")
	logger.NewLogger(logPath, logPrefix)

	/********************************************************************/
	/* CONFIGURE ECHO													*/
	/********************************************************************/
	timeout := time.Duration(viper.GetInt("Context.Timeout")) * time.Second

	e := echo.New()

	// Set timeout and disable keep alive
	e.Server.SetKeepAlivesEnabled(false)
	e.Server.ReadTimeout = timeout
	e.Server.WriteTimeout = timeout

	e.Use(middleware.CORS())

	// Init default cache
	initialize.Initialize()

	/********************************************************************/
	/* CONFIGURE Redis      											*/
	/********************************************************************/
	host := os.Getenv("REDIS_URL")
	fmt.Println(host)
	if host != "" {
		poolSize := viper.GetInt("Redis.PoolSize")
		minIdleConns := viper.GetInt("Redis.MinIdleConns")
		dB := viper.GetInt("Redis.DB")

		cacheManager := cache.CacheManager{}
		cacheManager.Init(host, poolSize, minIdleConns, dB)
		fmt.Println("Connect REDIS SUCCESSFULL!!!")
		account.Initialize(e, cacheManager, timeout)

	}

	homepage.Initialize(e)
	e.Start(viper.GetString("Server.Address"))

}
