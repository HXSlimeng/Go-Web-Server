package dto

import "github.com/HXSlimeng/Go-Web-Server/common"

type GetListDto struct {
	Leader string `form:"leader" json:"leader" search:"type:contains"` //负责人
	Phone  string `form:"phone" json:"phone"`                          //手机
	Email  string `form:"email" json:"email"`                          //邮箱
	Status int    `form:"status" json:"status"`                        //状态
	common.PageDto `gorm:"-"`
	common.Dto `gorm:"-"`
	
}

type DelIdsDto struct {
	Ids []int `form:"ids" json:"ids"`
}
