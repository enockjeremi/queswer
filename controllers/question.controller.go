package controllers

import (
	"net/http"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

type AnswerReponse struct {
	Content    string `json:"content"`
	QuestionID uint   `json:"question"`
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
		var answerReponse []AnswerReponse
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

func GetOneQuestion(c *gin.Context) {}
func PutQuestion(c *gin.Context)    {}
func DeleteQuestion(c *gin.Context) {}
