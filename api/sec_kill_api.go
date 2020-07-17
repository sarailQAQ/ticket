package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"server/message_queue"
	"server/middle_ware"
	"server/model"
	"server/resps"
	"strconv"
	"time"
)

func Sec_kill(c *gin.Context)  {
	t,ok := c.Get("user")
	if !ok {
		resps.Error(c,10001,errors.New("Unlogin"))
		return
	}
	user,_ := t.(middle_ware.UserClaim)
	film := c.PostForm("film")
	fid, _ := strconv.Atoi(film)
	ord := model.Order{
		Model:gorm.Model{
			CreatedAt: time.Now(),
		},
		Uid:   user.Id,
		Fid:   uint(fid),
	}
	err := message_queue.NewOrder(ord)
	if err != nil {
		log.Println(err)
		resps.Error(c,1001,errors.New("order failed"))
		return
	}
	resps.Ok(c)
}
