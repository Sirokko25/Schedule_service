package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"sheduler/internal/errorlist"
	"sheduler/internal/helpers"
)

func (h *Handlers) GetShedules(c *gin.Context) {
	user_id := c.Query("user_id")

	returnedShedules, err := h.SheduleStorage.FindShedules(user_id)
	if err != nil {
		log.Print(errorlist.ErrorFindShedule)
		c.JSON(http.StatusNotFound, string(err.Error()))
		return
	}
	log.Print("Records were successfully received")
	responseString := helpers.CreateResponceString(returnedShedules)
	c.IndentedJSON(http.StatusOK, responseString)
}
