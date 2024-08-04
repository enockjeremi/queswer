package http

import (
	"fmt"
	"net/http"

	"github.com/enockjeremi/queswer/delivery/http/middlewares"
	"github.com/enockjeremi/queswer/domain"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type QuestionHandler struct {
	QuestionUsecase domain.QuestionUsecase
}

func NewQuestionHandler(r *gin.Engine, qu domain.QuestionUsecase) {
	handler := &QuestionHandler{
		QuestionUsecase: qu,
	}

	v1 := r.Group("/v1")

	v1.GET("/question", handler.GetAllQuestion)
	v1.POST("/question", middlewares.CheckAuth, handler.CreateQuestion)
	v1.GET("/question/:id", handler.GetQuestion)
	v1.PUT("/question/:id", middlewares.CheckAuth, handler.UpdateQuestion)
	v1.DELETE("/question/:id", middlewares.CheckAuth, handler.RemoveQuestion)
	v1.PATCH("/question/:id/like", middlewares.CheckAuth, handler.LikeQuestion)
}

func (h *QuestionHandler) GetAllQuestion(c *gin.Context) {

	question, err := h.QuestionUsecase.GetAllQuestion()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &question,
	})
}

func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	var question domain.Question

	err := h.QuestionUsecase.GetQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "question not found")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	var question domain.Question
	currentUser, ok := c.Get("currentUser")
	fmt.Println(ok)
	user := currentUser.(domain.User)

	question.Answer = make([]domain.Answer, 0)
	question.User = user

	if err := c.ShouldBindBodyWithJSON(&question); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	err := h.QuestionUsecase.CreateQuestion(&question)
	if err != nil {
		utils.BadRequestResponse(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, Response{
			Success: true,
			Data:    &question,
		})
	}

}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	var question domain.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(domain.User).ID

	err := h.QuestionUsecase.GetQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "question not found")
		return
	}

	if question.UserID != userID {
		utils.ForbiddenResponse(c, "operation not allowed")
		return
	}
	c.BindJSON(&question)

	err = h.QuestionUsecase.UpdateQuestion(&question)
	if err != nil {
		utils.NotFoundResponse(c, "could not update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}

}

func (h *QuestionHandler) RemoveQuestion(c *gin.Context) {
	var question domain.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(domain.User).ID

	err := h.QuestionUsecase.GetQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "question not found")
		return
	}

	if question.UserID != userID {
		utils.ForbiddenResponse(c, "operation not allowed")
		return
	}

	err = h.QuestionUsecase.RemoveQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "could not delete")
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &question,
	})

}

func (h *QuestionHandler) LikeQuestion(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	questionId := c.Params.ByName("id")
	userId := currentUser.(domain.User).ID
	var question domain.Question

	err := h.QuestionUsecase.GetQuestion(&question, questionId)
	if err != nil {
		utils.NotFoundResponse(c, "question not found")
		return
	}

	like := utils.ContainsInSlice(question.Likes, int64(userId))
	if !like {
		question.Likes = append(question.Likes, int64(userId))
	} else {
		question.Likes = utils.RemoveElement(question.Likes, int64(userId))
	}

	err = h.QuestionUsecase.UpdateQuestion(&question)
	if err != nil {
		utils.NotFoundResponse(c, "could not update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}
}
