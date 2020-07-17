package api

import (
	"github.com/gin-gonic/gin"
	"server/middle_ware"
)

func Set_Router() {
	r := gin.Default()
	r.POST("/login",Login)
	r.POST("/register",Register)

	v := r.Group("/movie")
	{
		v.Use(middle_ware.LoginStatus)
		v.POST("/sec_kill",Sec_kill)
		v.POST("/new",NewTicket)
		}
	r.Run(":8080")
}
