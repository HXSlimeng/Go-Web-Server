package dto

import "github.com/HXSlimeng/Go-Web-Server/common"

type GetListDto struct {
	common.Dto
	Leader string `json:"leader" serach:"type:contains"`  //负责人
	Phone  string `json:"phone"`  //手机
	Email  string `json:"email"`  //邮箱
	Status int    `json:"status"` //状态
	Page   int    `json:"page"`
	PageNo int    `json:"pageNo"`
}