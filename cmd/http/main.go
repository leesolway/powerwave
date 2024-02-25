package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	httpadapters "github.com/leesolway/powerwave/internal/adapters/handlers/http"
	"github.com/leesolway/powerwave/internal/adapters/handlers/http/middleware"
	"github.com/leesolway/powerwave/internal/core/config"
	"github.com/leesolway/powerwave/internal/core/domain"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	powerMeterService := &domain.DefaultPowerMeterService{}

	router := SetupRouter(powerMeterService)
	router.Run(fmt.Sprintf(":%d", config.Port))
}

// SetupRouter creates a new gin router with all the necessary middleware and routes.
func SetupRouter(service domain.PowerMeterService) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMIddleware())

	RegisterRoutes(router, service)

	return router
}

// RegisterRoutes sets up all the route handlers for the application.
func RegisterRoutes(router *gin.Engine, service domain.PowerMeterService) {
	router.GET("/meters/:customer", httpadapters.AdapterForMetersByCustomer(service))
	router.GET("/readings/:serialID/:date", httpadapters.AdapterForMeterReading(service))
}
