package sysdept

import (
	"github.com/HXSlimeng/Go-Web-Server/app/sys_dept/dto"
	"github.com/HXSlimeng/Go-Web-Server/common"
	"github.com/HXSlimeng/Go-Web-Server/database"
	"gorm.io/gorm"
)

var SC *Service

func init()  {
	SC = new(Service)
}

type Service struct{}

func (s *Service) GetList(c *dto.GetListDto,list *[]Model,count *int64)error  {
    db := database.DB

	return db.Model(&Model{}).Scopes(
		c.Paginate,
		common.ResolveSearchQuery(c),
	).Count(count).Find(list).Error

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

func (s *Service)DelDept(ids *dto.DelIdsDto)error{
	db := database.DB.Delete(&Model{},ids.Ids)
	if db.Error !=nil {
		return db.Error	
	}
	return nil
}

type ExistErr struct{}

func (e *ExistErr)Error()string{
	return `data has exist!`
}
