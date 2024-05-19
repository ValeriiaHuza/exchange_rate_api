package main

import (
	"bytes"
	"encoding/json"
	"github.com/ValeriiaHuza/exchange_rate_api/controllers"
	"github.com/ValeriiaHuza/exchange_rate_api/initializers"
	"github.com/ValeriiaHuza/exchange_rate_api/migration"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	migration.AutomatedMigration()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	router.GET("/rate", controllers.ShowExchangeRate)
	router.POST("/subscribe", controllers.Subscribe)

	return router
}

func TestGetExchangeRate(t *testing.T) {

	router := router()

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/rate", nil)
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)

	var response float64
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestSubscribe(t *testing.T) {
	defer func() {
		initializers.DB.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	router := router()

	data := map[string]string{"email": "test@example.com"}
	jsonData, _ := json.Marshal(data)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response = w.Body.String()

	assert.Contains(t, response, "Email added")
}

func TestSubscribeInvalidEmail(t *testing.T) {

	router := router()

	data := map[string]string{"email": "invalid-data"}
	jsonData, _ := json.Marshal(data)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "error")
}
