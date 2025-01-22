package handlers

import (
	"net/http"
	"strconv"

	"survey-app/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AnswerHandler struct {
	db *gorm.DB
}

func NewAnswerHandler(db *gorm.DB) *AnswerHandler {
	return &AnswerHandler{db: db}
}

type CreateAnswerInput struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
}

func (h *AnswerHandler) CreateAnswer(c *gin.Context) {
	var input CreateAnswerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer := models.Answer{
		QuestionID: input.QuestionID,
		Text:       input.Text,
	}

	if err := h.db.Create(&answer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create answer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Answer created successfully", "answer_id": answer.ID})
}

func (h *AnswerHandler) ListAnswers(c *gin.Context) {
	var answers []models.Answer

	if err := h.db.Find(&answers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch answers"})
		return
	}

	c.JSON(http.StatusOK, answers)
}

func (h *AnswerHandler) GetAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer id"})
		return
	}

	var answer models.Answer
	if err := h.db.First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

type UpdateAnswerInput struct {
	Text string `json:"text" binding:"required"`
}

func (h *AnswerHandler) UpdateAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer id"})
		return
	}

	var input UpdateAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var answer models.Answer
	if err := h.db.First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	answer.Text = input.Text

	if err := h.db.Save(&answer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update answer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer updated successfully"})
}
func (h *AnswerHandler) DeleteAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer id"})
		return
	}

	var answer models.Answer
	if err := h.db.First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	if err := h.db.Delete(&answer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete answer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer deleted successfully"})
}
