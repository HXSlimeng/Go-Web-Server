package sysdept

import (
	"github.com/HXSlimeng/Go-Web-Server/app/sys_dept/dto"
	"github.com/HXSlimeng/Go-Web-Server/common"
	"github.com/gin-gonic/gin"
)



func GetList(c *gin.Context) {
	api := common.Api{Ctx:c}
	
	var (params dto.GetListDto;
	 list  = make([]Model,0);
	 count int64 = 0) 

	err := c.ShouldBindQuery(&params)

	if err != nil {
	 api.Error(err.Error())
	 return
	}

	SC.GetList(&params,&list,&count)

	api.PageOk(&count,&list,"请求成功！" )
}

func AddDept(c *gin.Context)  {
	api := common.Api{Ctx:c}

	var addTar Model
	
	err := c.Bind(&addTar)
	if err!=nil {
		api.Error(err.Error())
	}

	result := SC.AddDept(&addTar)
	if result != nil{
		 api.Error(result.Error())
			return
		}
	api.Ok("添加成功！",nil)
}

func DelDept(c *gin.Context){
	api := common.Api{Ctx:c}
	var params dto.DelIdsDto

	err := c.ShouldBind(&params)
	if err != nil {
		api.Error(err.Error())
		return
	}
	if err:= SC.DelDept(&params);err == nil{
		api.Ok("删除成功",nil)
	}else{
		api.Error(err.Error())
	}
}

func PatDept(c *gin.Context)  {
	
}