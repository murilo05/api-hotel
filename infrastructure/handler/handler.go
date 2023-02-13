package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
	errorUtils "gitlab.engdb.com.br/apigin/utils/error"
)

type ScheduleHandler struct {
	scheduleUseCase interfaces.ScheduleUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewProjectHandler(r *gin.Engine, us interfaces.ScheduleUseCase) *gin.Context {
	handler := &ScheduleHandler{
		scheduleUseCase: us,
	}
	r.GET("/agendas:disponibilidade", handler.checkScheduleAvailability)
	r.POST("/agendas", handler.createSchedule)
	r.GET("/agendas", handler.getSchedules)

	return nil
}

func (uh *ScheduleHandler) getSchedules(c *gin.Context) {
	ctx := context.Background()

	resp, code, errorResp := uh.scheduleUseCase.GetSchedules(ctx)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(code, resp)
}

func (uh *ScheduleHandler) createSchedule(c *gin.Context) {
	ctx := context.Background()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errResp := errorUtils.CreateError(500, err.Error())
		c.JSON(500, errResp)
		return
	}

	body := entities.InputSchedule{}

	json.Unmarshal(jsonData, &body)
	if err != nil {
		errResp := errorUtils.CreateError(500, err.Error())
		c.JSON(500, errResp)
		return
	}

	code, errorResp := uh.scheduleUseCase.CreateSchedule(ctx, body)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(200, gin.H{"Data": "Agendamento Criado com sucesso"})
}

func (uh *ScheduleHandler) checkScheduleAvailability(c *gin.Context) {
	ctx := context.Background()

	resp, code, errorResp := uh.scheduleUseCase.CheckScheduleAvailability(ctx)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(200, resp)
}
