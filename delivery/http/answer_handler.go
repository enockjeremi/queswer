package http

import (
	"net/http"
	"strconv"

	"github.com/enockjeremi/queswer/delivery/http/middlewares"
	"github.com/enockjeremi/queswer/domain"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
)

type AnswerHandler struct {
	AnswerUsecase domain.AnswerUsecase
}

func NewAnswerHandler(r *gin.Engine, au domain.AnswerUsecase) {
	handler := &AnswerHandler{
		AnswerUsecase: au,
	}

	v1 := r.Group("/v1")

	v1.GET("answer", handler.GetAllAnswer)
	v1.GET("answer/:id", handler.GetOneAnswer)
	v1.POST("answer", middlewares.CheckAuth, handler.CreateAnswer)
	v1.PUT("answer/:id", middlewares.CheckAuth, handler.UpdateAnswer)
	v1.DELETE("answer/:id", middlewares.CheckAuth, handler.RemoveAnswer)
}

func (h *AnswerHandler) GetAllAnswer(c *gin.Context) {

	answer, err := h.AnswerUsecase.GetAllAnswer()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answer,
	})
}

func toString(i uint) string {
	v := uint64(i)
	return strconv.FormatUint(uint64(v), 10)
}

func (h *AnswerHandler) CreateAnswer(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(domain.User)

	var answer domain.Answer
	var question domain.Question

	answer.User = user
	if err := c.ShouldBindBodyWithJSON(&answer); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}
	questionID := toString(answer.QuestionID)

	err := h.AnswerUsecase.GetQuestion(&question, questionID)
	if err != nil {
		utils.NotFoundResponse(c, "question not found")
		return
	} else {
		err := h.AnswerUsecase.CreateAnswer(&answer)
		if err != nil {
			utils.BadRequestResponse(c, err)
			return
		} else {
			c.JSON(http.StatusCreated, Response{
				Success: true,
				Data:    &answer,
			})
		}
	}

}

func (h *AnswerHandler) GetOneAnswer(c *gin.Context) {
	var answer domain.Answer
	id := c.Params.ByName("id")
	err := h.AnswerUsecase.GetAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "answer not found")
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answer,
	})
}

func (h *AnswerHandler) UpdateAnswer(c *gin.Context) {
	var answer domain.Answer
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(domain.User).ID

	err := h.AnswerUsecase.GetAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "answer not found")
		return
	}
	c.BindJSON(&answer)

	if answer.UserID != userID {
		utils.ForbiddenResponse(c, "operation not allowed")
		return
	}

	err = h.AnswerUsecase.UpdateAnswer(&answer)
	if err != nil {
		utils.NotFoundResponse(c, "could not update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &answer,
		})
	}
}

func (h *AnswerHandler) RemoveAnswer(c *gin.Context) {
	var answer domain.Answer
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(domain.User).ID

	err := h.AnswerUsecase.GetAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "answer not found")
		return
	}

	if answer.UserID != userID {
		utils.ForbiddenResponse(c, "operation not allowed")
		return
	}

	err = h.AnswerUsecase.RemoveAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "could not delete")
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answer,
	})

}
