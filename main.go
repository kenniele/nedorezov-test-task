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
	defer connection.Close()

	router := gin.Default()
	router.Static("/assets/", "front/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", handlerIndex)
	router.POST("/accounts", handlerAccountRegistration)
	router.POST("/accounts/:id/deposit", handlerDeposit)
	router.POST("/accounts/:id/withdraw", handlerWithdraw)
	router.POST("/accounts/:id/balance", HandlerBalance)
	log.Println("Running server on port 8080")
	if err = router.Run(":8080"); err != nil {
		log.Fatal("Error while start running server", err)
		return
	}
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
		log.Println("Error while binding JSON", err)
	}

	if id, err = account.Create(); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"Error": err,
		})
		log.Println("Error while creating an account", err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"Error": nil,
		"ID":    id,
	})

}

func handlerDeposit(c *gin.Context) {
	var operation Operation
	var account Account
	if err := c.BindJSON(&operation); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"Error": err.Error(),
		})
		log.Println("Error while binding JSON", err)
		return
	}

	resultChannel := make(chan Result)

	go func() {
		if err := account.Select(operation); err != nil {
			resultChannel <- Result{Error: err}
			return
		}

		if err := account.Deposit(operation.Amount); err != nil {
			resultChannel <- Result{Error: err}
			log.Println("Error while depositing account", err)
			return
		}

		resultChannel <- Result{ID: account.ID, Balance: account.Balance, Error: nil}
	}()
	result := <-resultChannel

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": result.Error.Error(),
		})
		log.Println("Error while processing deposit", result.Error)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Error":   nil,
		"ID":      result.ID,
		"Balance": operation.Amount,
	})
}

func handlerWithdraw(c *gin.Context) {
	var operation Operation
	var account Account

	if err := c.BindJSON(&operation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resultChannel := make(chan Result)

	go func() {
		if err := account.Select(operation); err != nil {
			resultChannel <- Result{Error: err}
			return
		}

		if err := account.Withdraw(operation.Amount); err != nil {
			resultChannel <- Result{Error: err}
			return
		}

		resultChannel <- Result{ID: account.ID, Balance: account.Balance, Error: nil}
	}()

	result := <-resultChannel

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": result.Error.Error(),
		})
		log.Println("Error while processing withdraw", result.Error)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Error":   nil,
		"ID":      result.ID,
		"Balance": result.Balance,
	})
}

func HandlerBalance(c *gin.Context) {
	var account Account
	if err := c.BindJSON(&account); err != nil {
		log.Println("Error while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := account.isIn(); err != nil {
		log.Println("Error while checking account", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resultChannel := make(chan Result)

	go func() {
		resultChannel <- account.GetBalance()
	}()

	result := <-resultChannel

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": result.Error.Error(),
		})
		log.Println("Error while getting balance", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Error":   nil,
		"ID":      account.ID,
		"Balance": result.Balance,
	})
}
