package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"nedorezov-test-task/config"
	"net/http"
)

var connection *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var err error
	conf := config.New()
	log.Println("Connecting to config file")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.DB.HOST, conf.DB.PORT, conf.DB.USER, conf.DB.PASSWORD, conf.DB.DBNAME, conf.DB.SSLMODE)

	if connection, err = sql.Open("postgres", connectionString); err != nil {
		log.Fatal(err)
	}
	log.Println("Connecting to database")

	router := gin.Default()
	router.Static("/assets/", "front/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", handlerIndex)
	router.POST("/accounts", handlerAccountRegistration)
	log.Println("Running server on port 8080")
	_ = router.Run(":8080")
}

func handlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Error": nil,
	})
}

func handlerAccountRegistration(c *gin.Context) {
	var account Account
	var id int
	var err error

	if err = c.BindJSON(&account); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"Error": err.Error(),
		})
		log.Fatal(err)
	}

	if id, err = account.Create(); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"Error": err,
		})
		log.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"Error": nil,
		"ID":    id,
	})

}
