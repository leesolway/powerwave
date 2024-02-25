package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	httpadapters "github.com/leesolway/powerwave/internal/adapters/handlers/http"
	"github.com/leesolway/powerwave/internal/adapters/handlers/http/middleware"
	"github.com/leesolway/powerwave/internal/core/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	router := SetupRouter()
	router.Run(fmt.Sprintf(":%d", config.Port))
}

// SetupRouter creates a new gin router with all the necessary middleware and routes.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMIddleware())

	RegisterRoutes(router)

	return router
}

// RegisterRoutes sets up all the route handlers for the application.
func RegisterRoutes(router *gin.Engine) {
	router.GET("/meters/:customer", httpadapters.AdapterForMetersByCustomer)
	router.GET("/readings/:serialID/:date", httpadapters.AdapterForMeterReading)
}
