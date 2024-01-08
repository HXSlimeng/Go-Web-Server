package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Ctx *gin.Context
}


func (a *Api)Ok(obj interface {},message string){
	a.Ctx.JSON(http.StatusOK,Response{http.StatusOK,message,obj})
}

func (a *Api)Error(message string){
	errorCode := http.StatusInternalServerError
	a.Ctx.JSON(errorCode,Response{errorCode,message,nil})
}

func (a *Api)MakeContext(c *gin.Context) *Api {
	a.Ctx = c
	return a
}

func (a *Api)PageOk(page int ,pageSize int ,total int ,data interface {},message string){
	res := Page{total,page,pageSize,data}
	a.Ok(res,message)
}