package sysdept

import (
	"github.com/HXSlimeng/Go-Web-Server/app/sys_dept/dto"
	"github.com/HXSlimeng/Go-Web-Server/database"
)

type Service struct{}


func (s *Service) GetList(req *dto.GetListDto) []Model {
	db := database.DB
	
	var users []Model
	req.ResolveSearchQuery(db)
	db.Find(&users).Limit(10).Offset( 1 * 10)
	return users
}