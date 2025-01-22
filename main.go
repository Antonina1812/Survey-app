package main

import (
	"log"
	"survey-app/internal/database"
	"survey-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.SetupDB()

	r := gin.Default()

	userHandler := handlers.NewUserHandler(db)
	r.POST("/register", userHandler.RegisterUser)
	r.POST("/login", userHandler.LoginUser)

	pollHandler := handlers.NewPollHandler(db)
	r.POST("/polls", pollHandler.CreatePoll)
	r.GET("/polls", pollHandler.ListPolls)
	r.GET("/polls/:id", pollHandler.GetPoll)
	r.PUT("/polls/:id", pollHandler.UpdatePoll)
	r.DELETE("/polls/:id", pollHandler.DeletePoll)

	questionHandler := handlers.NewQuestionHandler(db)
	r.POST("/questions", questionHandler.CreateQuestion)
	r.GET("/questions", questionHandler.ListQuestions)
	r.GET("/questions/:id", questionHandler.GetQuestion)
	r.PUT("/questions/:id", questionHandler.UpdateQuestion)
	r.DELETE("/questions/:id", questionHandler.DeleteQuestion)

	answerHandler := handlers.NewAnswerHandler(db)
	r.POST("/answers", answerHandler.CreateAnswer)
	r.GET("/answers", answerHandler.ListAnswers)
	r.GET("/answers/:id", answerHandler.GetAnswer)
	r.PUT("/answers/:id", answerHandler.UpdateAnswer)
	r.DELETE("/answers/:id", answerHandler.DeleteAnswer)

	responseHandler := handlers.NewResponseHandler(db)
	r.POST("/responses", responseHandler.CreateResponse)
	r.GET("/responses", responseHandler.ListResponses)
	r.GET("/responses/:id", responseHandler.GetResponse)

	r.POST("/responseanswers", responseHandler.CreateResponseAnswer)

	log.Fatal(r.Run(":8080"))
}

//go run main.go and visit 0.0.0.0:8080/register on browser
