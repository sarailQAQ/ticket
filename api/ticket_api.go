package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"server/model"
	"server/resps"
	"time"
)

type Ticket struct {
	Id uint `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Begin time.Time `json:"begin" binding:"required"`
	End time.Time `json:"end" binding:"required"`
	Site string `json:"site" binding:"required"`
	Total int `json:"total" binding:"required"`
	Left int `json:"left" binding:"required"`
}

func NewTicket(c *gin.Context) {
	var data Ticket
	if err  := c.ShouldBindJSON(&data); err != nil {
		resps.Error(c,1001,err)
		return
	}
	ticket := model.Ticket{
		Model: gorm.Model{ID:data.Id},
		Name:  data.Name,
		Begin: data.Begin,
		End:   data.End,
		Site:  data.Site,
		Total: data.Total,
		Left:  data.Left,
	}
	err := model.NewTicket(ticket)
	if err != nil {
		resps.Error(c,1001,err)
		return
	}else {
		resps.Ok(c)
	}
}
