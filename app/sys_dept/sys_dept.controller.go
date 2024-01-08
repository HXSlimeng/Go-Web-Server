package sysdept

import (
	"github.com/HXSlimeng/Go-Web-Server/app/sys_dept/dto"
	"github.com/HXSlimeng/Go-Web-Server/common"
	"github.com/gin-gonic/gin"
)



func GetList(c *gin.Context) {
	api := common.Api{Ctx:c}

	params := &dto.GetListDto{}
	err := c.ShouldBindQuery(params)
	if err != nil {
	 api.Error(err.Error())
	 return
	}
	sysDeptService := new(Service)
	list := sysDeptService.GetList(params)
	api.Ok(list,"请求成功")
}