package routes

import (
	"example.com/rest_api/models"
	"example.com/rest_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // base คือเลขฐาน
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	evetns, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	//context.JSON(http.StatusOK, gin.H{ // H = map[string]any
	//	"message": "Hello!",
	//})

	context.JSON(http.StatusOK, evetns)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize!"})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // base คือเลขฐาน
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}

	updID := strconv.FormatInt(updatedEvent.ID, 10)
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully!",
		updID:     updatedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // base คือเลขฐาน
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully!",
	})
}
