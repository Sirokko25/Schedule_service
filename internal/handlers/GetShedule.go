package handlers

import (
	"log"
	"net/http"
	"sheduler/internal/errorlist"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetShedule(c *gin.Context) {
	user_id := c.Query("user_id")
	shedule_id := c.Query("shedule_id")

	shedule, err := h.SheduleStorage.FindShedule(user_id, shedule_id)
	if err != nil {
		log.Printf(errorlist.ErrorFindShedule)
		c.JSON(http.StatusNotFound, string(err.Error()))
		return
	}
	log.Print("The record were successfully received")
	c.IndentedJSON(http.StatusOK, shedule)
}
