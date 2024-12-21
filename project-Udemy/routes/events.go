package routes

import (
	"net/http"
	"strconv"

	"github.com/Wong-bui/Udemy-project/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "Could not get events"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Hello!", "events": events})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Get event success", "event": event})

}

func createEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "Could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("event_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not fetch event."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err, "message": "Not matching userId"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not parse request event."})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not update event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Update event success"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err, "message": "Not matching userId"})
		return
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not fetch event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Could not delete event."})
		return
	}

	context.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Event deleted successfully."})

}
