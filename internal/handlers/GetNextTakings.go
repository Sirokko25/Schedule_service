package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetNextTakings(c *gin.Context) {
	user_id := c.Query("user_id")
	time, err := h.SheduleStorage.NextTakings(user_id)
	if err != nil {
		log.Printf("Error in calculating the nearest medication intake time")
		c.JSON(http.StatusNotFound, string(err.Error()))
		return
	}
	log.Print("The record was successfully received")
	c.IndentedJSON(http.StatusOK, time)
}
