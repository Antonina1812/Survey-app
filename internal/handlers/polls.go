package handlers

import (
	"net/http"
	"strconv"

	"survey-app/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PollHandler struct {
	db *gorm.DB
}

func NewPollHandler(db *gorm.DB) *PollHandler {
	return &PollHandler{db: db}
}

type CreatePollInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	OwnerID     uint   `json:"owner_id" binding:"required"`
}

func (h *PollHandler) CreatePoll(c *gin.Context) {
	var input CreatePollInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poll := models.Poll{
		Title:       input.Title,
		Description: input.Description,
		OwnerID:     input.OwnerID,
	}

	if err := h.db.Create(&poll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create poll"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Poll created successfully", "poll_id": poll.ID})
}

func (h *PollHandler) ListPolls(c *gin.Context) {
	var polls []models.Poll

	if err := h.db.Find(&polls).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch polls"})
		return
	}

	c.JSON(http.StatusOK, polls)
}

func (h *PollHandler) GetPoll(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid poll id"})
		return
	}

	var poll models.Poll
	if err := h.db.First(&poll, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}

	c.JSON(http.StatusOK, poll)
}

type UpdatePollInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *PollHandler) UpdatePoll(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid poll id"})
		return
	}

	var input UpdatePollInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var poll models.Poll
	if err := h.db.First(&poll, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}

	poll.Title = input.Title
	poll.Description = input.Description

	if err := h.db.Save(&poll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update poll"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Poll updated successfully"})
}

func (h *PollHandler) DeletePoll(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid poll id"})
		return
	}

	var poll models.Poll
	if err := h.db.First(&poll, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}

	if err := h.db.Delete(&poll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete poll"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Poll deleted successfully"})
}
