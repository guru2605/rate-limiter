package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
)

const REQUEST_LIMIT_WINDOW = 10 // seconds
const REQUEST_LIMIT = 10

type json map[string]interface{}

type ILimiterMiddleWare interface {
	IsRequestValid() gin.HandlerFunc
	RunGoFuncToResetLimit()
}

// Limits
type Limiter struct {
	currentRequestCount int
	requestLimit        int
}

var limiterInstance Limiter

func New() ILimiterMiddleWare {
	if limiterInstance.requestLimit == 0 {
		limiterInstance := Limiter{
			currentRequestCount: 0,
			requestLimit:        REQUEST_LIMIT,
		}
		return &limiterInstance
	}
	return &limiterInstance
}

func (l *Limiter) IsRequestValid() gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		l.currentRequestCount++

		if l.currentRequestCount < l.requestLimit {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(400, json{
				"message": "Not Allowed",
			})
		}
	}

}

func (l *Limiter) RunGoFuncToResetLimit() {
	go func() {
		for {
			l.currentRequestCount = 0
			time.Sleep(time.Duration(time.Second * REQUEST_LIMIT_WINDOW))
		}
	}()

}
