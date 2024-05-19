package main

import (
	"fmt"
	"github.com/ValeriiaHuza/exchange_rate_api/controllers"
	"github.com/ValeriiaHuza/exchange_rate_api/initializers"
	"github.com/ValeriiaHuza/exchange_rate_api/migration"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	migration.AutomatedMigration()
}

//TODO: write comments

func main() {

	router := gin.Default()

	router.GET("/rate", controllers.ShowExchangeRate)
	router.POST("/subscribe", controllers.Subscribe)

	c := cron.New()
	_, err := c.AddFunc("0 9 * * *", controllers.SendEmails) // Every day at 9:00 AM
	if err != nil {
		fmt.Println("Error with scheduler : ", err)
	}
	c.Start()

	err = router.Run(":8000")
	if err != nil {
		fmt.Println("Error with starting server:", err)
		return
	}

}
