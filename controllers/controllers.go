package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ValeriiaHuza/exchange_rate_api/initializers"
	"github.com/ValeriiaHuza/exchange_rate_api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func getExchangeRate() (float64, error) {
	response, err := http.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/dollar_info?json")

	if err != nil {
		return 0, fmt.Errorf("could not get exchange rate: %w", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("could not read response body: %w", err)
	}

	var dollarRate []struct {
		Rate float64 `json:"rate"`
	}

	if err := json.Unmarshal(body, &dollarRate); err != nil {
		return 0, fmt.Errorf("could not parse json response: %w", err)
	}

	if len(dollarRate) == 0 {
		return 0, fmt.Errorf("no rates found: %w", err)
	}

	return dollarRate[0].Rate, nil
}

func ShowExchangeRate(c *gin.Context) {

	rate, err := getExchangeRate()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rate)
}

func Subscribe(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with data binding"})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email can not be empty"})
		return
	}

	if !isValidEmail(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	var existingUser models.User

	if err := initializers.DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already subscribed"})
		return
	}

	user := models.User{Email: body.Email}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}

	c.JSON(http.StatusOK, "Email added : "+user.Email)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func SendEmails() {

	var emails []models.User

	initializers.DB.Find(&emails)

	rate, err := getExchangeRate()

	if err != nil {
		fmt.Println("Failed to get exchange rate", err)
		return
	}

	sendEmail := os.Getenv("MAIL_EMAIL")
	sendPassword := os.Getenv("MAIL_PASSWORD")

	for _, user := range emails {
		m := gomail.NewMessage()
		m.SetHeader("From", sendEmail)
		m.SetHeader("To", user.Email)
		m.SetHeader("Subject", "Daily USD/UAH Rate")
		m.SetBody("text/plain", fmt.Sprintf("The current rate is: %f", rate))

		d := gomail.NewDialer("smtp.gmail.com", 587, sendEmail, sendPassword)

		if err := d.DialAndSend(m); err != nil {
			log.Println("Failed to send email to", user.Email, ":", err)
		}

		fmt.Println("Email sent to - ", user.Email)
	}

}
