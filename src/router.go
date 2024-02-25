package main

import (
	httpadapters "github.com/leesolway/powerwave/src/adapters/http"
	"github.com/leesolway/powerwave/src/middleware"

	"github.com/gin-gonic/gin"
)

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
