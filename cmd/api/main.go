package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/handlers"
	"github.com/imabg/authn/store"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	conn, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Authn is up and running",
		})
	})
	// source store
	sourceStore := store.NewSourceStore(conn)
	sourceHandler := handlers.NewSourceHandler(sourceStore)
	var base = "/api/v1/source"
	router.POST(base, sourceHandler.Create)
	router.GET(fmt.Sprintf("%s/:id", base), sourceHandler.GetByID)
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
