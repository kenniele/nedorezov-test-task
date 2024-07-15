package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"nedorezov-test-task/config"
)

var connection *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	fmt.Println(conf.HOST)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.DB.HOST, conf.DB.PORT, conf.DB.USER, conf.DB.PASSWORD, conf.DB.DBNAME, conf.DB.SSLMODE)

	if connection, err := sql.Open("postgres", connectionString); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", handlerIndex)

	_ = router.Run(":8080")
}

func handlerIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"Error": nil,
	})
}
