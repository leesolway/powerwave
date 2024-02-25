package httpadapters

import (
	"net/http"

	"github.com/leesolway/powerwave/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// AdapterForMetersByCustomer adapts HTTP requests to domain logic for fetching meters by customer
func AdapterForMetersByCustomer(service domain.PowerMeterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer := c.Param("customer")

		meters, err := service.GetMetersByCustomerName(customer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch meters"})
			return
		}

		if len(meters) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "No meters found"})
			return
		}

		c.JSON(http.StatusOK, meters)
	}
}
