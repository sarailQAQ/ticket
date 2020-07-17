package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"server/jwt"
	"server/model"
	"server/resps"
)

// user.go something

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var f LoginForm
	if err := c.ShouldBindJSON(&f); err != nil {
		resps.FormError(c)
		return
	}

	id := model.Login(f.Username,f.Password)
	if id != 0 {
		token := jwt.Creat(f.Username,id)
		resps.OkWithData(c,gin.H{
			"token":token,
			"username": f.Username,
		})	// login success
	}else {

		resps.Error(c,1001,errors.New("password error"))
	}
}

func Register(c *gin.Context) {
	var f LoginForm
	if err := c.ShouldBindJSON(&f); err != nil {
		resps.FormError(c)
		return
	}

	err := model.Register(f.Username,f.Password)
	if err != nil {
		resps.Error(c,1001,err)
	}else {
		resps.OkWithData(c, gin.H{})
	}
}



