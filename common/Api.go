package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Ctx *gin.Context
}


func (a *Api)Ok(message string,data interface{}){
	a.Ctx.JSON(http.StatusOK,Response{http.StatusOK,message,data})
}

func (a *Api)Error(message string){
	errorCode := http.StatusInternalServerError
	a.Ctx.JSON(errorCode,Response{errorCode,message,nil})
}

func (a *Api)MakeContext(c *gin.Context) *Api {
	a.Ctx = c
	return a
}

func (a *Api)PageOk(pageInfo PageDto,total *int64 ,data interface {},message string){
	res := PageRes{Response: Response{http.StatusOK,message,data},Total: *total,Page: pageInfo.Page,PageSize: pageInfo.PageSize}
	a.Ctx.JSON(http.StatusOK,res)
}