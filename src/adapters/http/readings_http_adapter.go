package httpadapters

import (
	"net/http"
	"time"

	"github.com/leesolway/powerwave/src/domain"

	"github.com/gin-gonic/gin"
)

// AdapterForMeterReading adapts HTTP requests to domain logic for fetching meter readings by date
func AdapterForMeterReading(c *gin.Context) {
	serialID := c.Param("serialID")
	date, err := time.Parse("2006-01-02", c.Param("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	reading, err := domain.GetMeterReadingBySerialIDAndDate(serialID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reading"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"serialID": serialID, "reading": reading})
}
