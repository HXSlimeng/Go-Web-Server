package sysdept

import (
	"github.com/HXSlimeng/Go-Web-Server/app/sys_dept/dto"
	"github.com/HXSlimeng/Go-Web-Server/common"
	"github.com/gin-gonic/gin"
)



func GetList(c *gin.Context) {
	api := common.Api{Ctx:c}
	
	var params dto.GetListDto
	err := c.ShouldBindQuery(&params)

	if err != nil {
	 api.Error(err.Error())
	 return
	}

	sysDeptService := new(Service)
	list,count := sysDeptService.GetList(&params)

	api.PageOk(params.PageDto,count,list,"请求成功！" )
}

func AddDept(c *gin.Context)  {
	api := common.Api{Ctx:c}
	
	var addTar Model
	
	err := c.Bind(&addTar)
	if err!=nil {
		api.Error(err.Error())
	}

	service := new(Service)
	result := service.AddDept(&addTar)
	if result != nil{
		 api.Error(result.Error())
			return
		}
	api.Ok("添加成功！",nil)
}