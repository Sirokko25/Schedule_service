package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"sheduler/internal/helpers"
	"sheduler/internal/models"
)

func (h *Handlers) AddShedule(c *gin.Context) {
	var userShedule models.Shedule
	err := c.BindJSON(&userShedule)
	if err != nil {
		log.Fatalf("Incorrect data was entered in the request body: %v", err)
		c.JSON(http.StatusBadRequest, string(err.Error()))
		return
	}
	log.Printf("The request came with the data: userid = %d, medicine = %s, period = %d, duration = %s", userShedule.UserId, userShedule.Medicine, userShedule.Period, userShedule.Duration)
	ok := helpers.ValidateShedule(userShedule)
	if !ok {
		log.Print("The data in the request is incorrect")
		c.JSON(http.StatusBadRequest, "The data in the request is incorrect")
		return
	}

	id, err := h.SheduleStorage.AppendShedule(userShedule)
	if err != nil {
		log.Fatalf("Error when adding a task to the database: %v", err)
		c.JSON(http.StatusInternalServerError, string(err.Error()))
		return
	}
	log.Print("The entry was added successfully")
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("The entry with id = %d has been successfully added", id))
}
