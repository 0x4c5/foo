package bootstrap

import (
	"foo/facades"
	"foo/middleware"
	"foo/pkg/redis"
	"foo/pkg/vipper"
	"foo/pkg/zap"
	"foo/route"

	"github.com/gin-gonic/gin"
)

func Boot() (err error) {
	// load config
	facades.Viper, err = vipper.Init(".env")
	if err != nil {
		return
	}
	// load logger
	facades.Logger, err = zap.Init()
	if err != nil {
		return
	}
	// load redis-client
	facades.Redis, err = redis.Init(facades.Viper.GetString("REDIS_ADDR"), facades.Viper.GetString("REDIS_PASSWORD"), facades.Viper.GetInt("REDIS_DB"), facades.Viper.GetInt("REDIS_POOL_SIZE"))
	if err != nil {
		return
	}
	// register gin-handler
	if facades.Viper.GetBool("DEBUG") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	facades.Handler = gin.New()
	// register middlewares
	facades.Handler.Use(middleware.Recovery(true), middleware.Logger())
	// register routes
	route.RegisterAPI(facades.Handler.Group("/api/v1"))
	return nil
}
