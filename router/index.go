package router

import (
	sysdept "github.com/HXSlimeng/Go-Web-Server/app/sys_dept"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine {
	router := gin.Default()
	{
		deptR := router.Group("/dept")
		deptR.GET("/getList",sysdept.GetList)
		deptR.POST("/addDept",sysdept.AddDept)
		deptR.DELETE("/delDept",sysdept.DelDept)
		deptR.PATCH("patchDept",sysdept.PatDept)
	}
	return router
}