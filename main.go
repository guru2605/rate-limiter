package main

import (
	"fmt"
	"log"
	"os"
	"rate-limiter/limiter"

	_ "github.com/lib/pq" // Import the PostgreSQL driver

	"github.com/gin-gonic/gin"
)

type json map[string]interface{}

func home(pool *DBPool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, json{
			"message": "API Success",
		})
	}
}

const PG_CONNECTION_STRING = "user=postgres dbname=mydb password=mypassword host=localhost port=5432 sslmode=disable"

func main() {
	app := gin.Default()

	limiterOb := limiter.New()
	limiterOb.RunGoFuncToResetLimit()

	pool := New(10, PG_CONNECTION_STRING)

	app.GET("", limiterOb.IsRequestValid(), home(pool))

	fmt.Println("Successfully connected to the database!")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
		log.Printf("defaulting to port %s", port)
	}

	app.Run(":" + port)

}
