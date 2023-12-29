package main

import (
	"rate-limiter/limiter"

	"github.com/gin-gonic/gin"
)

type json map[string]interface{}

func home() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, json{
			"message": "API Success",
		})
	}
}

func main() {
	app := gin.Default()

	limiterOb := limiter.New()
	limiterOb.RunGoFuncToResetLimit()

	app.GET("", limiterOb.IsRequestValid(), home())

	app.Run(":5001")

}
