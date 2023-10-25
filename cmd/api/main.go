package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/db"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	fmt.Println("Connected to the database 🚀")
	r.Run()
}
