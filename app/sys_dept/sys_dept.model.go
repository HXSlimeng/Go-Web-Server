package sysdept

import (
	"github.com/gin-gonic/gin"
)

type Model struct {
	DeptId   int    `json:"deptId"`
	ParentId int    `json:"parentId"`
	DeptName string `json:"deptName"`
	DeptPath string `json:"deptPath"`
	Sort     int    `json:"sort"`
}

func (entity *Model) GetList(c *gin.Context) {
	
}

func (Model) TableName() string {
	return "sys_dept"
}