package controllers

import (
	"net/http"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

type AnswerReponse struct {
	Content string `json:"content"`
}

type QuestionReponse struct {
	ID          uint            `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Completed   bool            `json:"completed"`
	Answers     []AnswerReponse `json:"answers"`
}

func QuestionSerializer(q []models.Question) []QuestionReponse {
	var questionResponse []QuestionReponse
	for _, question := range q {
		var answerReponse = make([]AnswerReponse, 0)
		for _, answer := range question.Answer {
			answerReponse = append(answerReponse, AnswerReponse{
				Content: answer.Content,
			})
		}
		questionResponse = append(questionResponse, QuestionReponse{
			ID:          question.ID,
			Title:       question.Title,
			Description: question.Description,
			Completed:   question.Completed,
			Answers:     answerReponse,
		})
	}
	return questionResponse
}

func GetAllQuestion(c *gin.Context) {
	var question []models.Question
	if err := services.FindAllQuestion(&question); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	response := QuestionSerializer(question)

	c.JSON(http.StatusOK, &response)
}
func PostQuestion(c *gin.Context) {
	var question models.Question
	question.Answer = make([]models.Answer, 0)

	if err := c.ShouldBindBodyWithJSON(&question); err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateQuestion(&question)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, question)
	}
}

func GetOneQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	var question models.Question
	err := services.GetOneQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
	} else {
		c.JSON(http.StatusOK, question)
	}
}
func PutQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")
	err := services.GetOneQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
	}
	c.BindJSON(&question)
	err = services.UpdateQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
	} else {
		c.JSON(http.StatusOK, question)
	}

}
func DeleteQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")
	err := services.GetOneQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
	}
	c.BindJSON(&question)
	err = services.DeleteQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
	} else {
		c.JSON(http.StatusOK, question)
	}
}
