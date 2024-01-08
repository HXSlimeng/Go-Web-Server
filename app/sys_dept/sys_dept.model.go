package sysdept

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	DeptId   int    `json:"deptId" gorm:"primaryKey;autoIncrement;"` //部门编码
	ParentId int    `json:"parentId" gorm:""`                        //上级部门
	DeptName string `json:"deptName"  gorm:"size:128;"`              //部门名称
	Sort     int    `json:"sort" gorm:"size:4;"`                     //排序
	Leader   string `json:"leader" gorm:"size:128;"`                 //负责人
	Phone    string `json:"phone" gorm:"size:11;"`                   //手机
	Email    string `json:"email" gorm:"size:64;"`                   //邮箱
	Status   int    `json:"status" gorm:"size:4;"`                   //状态
	DataScope string    `json:"dataScope" gorm:"-"`
	Params    string    `json:"params" gorm:"-"`
	Children  []Model `json:"children" gorm:"-"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}



func (*Model) TableName() string {
	return "sys_dept"
}