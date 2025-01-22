package handlers

import (
	"net/http"
	"strconv"

	"survey-app/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ResponseHandler struct {
	db *gorm.DB
}

func NewResponseHandler(db *gorm.DB) *ResponseHandler {
	return &ResponseHandler{db: db}
}

type CreateResponseInput struct {
	PollID uint `json:"poll_id" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

func (h *ResponseHandler) CreateResponse(c *gin.Context) {
	var input CreateResponseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.Response{
		PollID: input.PollID,
		UserID: input.UserID,
	}

	if err := h.db.Create(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Response created successfully", "response_id": response.ID})
}
func (h *ResponseHandler) ListResponses(c *gin.Context) {
	var responses []models.Response

	if err := h.db.Find(&responses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch responses"})
		return
	}

	c.JSON(http.StatusOK, responses)
}

func (h *ResponseHandler) GetResponse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid response id"})
		return
	}

	var response models.Response
	if err := h.db.First(&response, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Response not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}

type CreateResponseAnswerInput struct {
	ResponseID uint `json:"response_id" binding:"required"`
	QuestionID uint `json:"question_id" binding:"required"`
	AnswerID   uint `json:"answer_id" binding:"required"`
}

func (h *ResponseHandler) CreateResponseAnswer(c *gin.Context) {
	var input CreateResponseAnswerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseAnswer := models.ResponseAnswer{
		ResponseID: input.ResponseID,
		QuestionID: input.QuestionID,
		AnswerID:   input.AnswerID,
	}

	if err := h.db.Create(&responseAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response answer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Response answer created successfully", "response_answer_id": responseAnswer.ID})
}
