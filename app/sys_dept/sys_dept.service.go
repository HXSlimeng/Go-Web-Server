package sysdept

import (
	"github.com/HXSlimeng/Go-Web-Server/app/sys_dept/dto"
	"github.com/HXSlimeng/Go-Web-Server/database"
	"gorm.io/gorm"
)

type Service struct{}


func (s *Service) GetList(req *dto.GetListDto) (*[]Model,*int64) {
    db := database.DB

    var users []Model
	var total int64

	db.Model(&Model{}).Count(&total)

	db = req.ResolveSearchQuery(req, db)
	db = req.ResolvePage(db)

	db.Find(&users)

    return &users, &total
}

func (s *Service)AddDept(dept *Model)error  {
	db := database.DB

	md := Model{}
	result := db.Where(dept).First(&md)
	
	if result.Error == gorm.ErrRecordNotFound {
		return db.Create(dept).Error
	}else{
		return &ExistErr{}
	}
}

type ExistErr struct{}

func (e *ExistErr)Error()string{
	return `data has exist!`
}