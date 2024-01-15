package common

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Dto struct {
}

func (dto *Dto)ResolveSearchQuery (ins interface{}, db *gorm.DB) *gorm.DB{

	dtoType := reflect.TypeOf(ins).Elem()
	dtoVal := reflect.ValueOf(ins).Elem()
	nums := dtoType.NumField()

	for i := 0; i < nums; i++ {
		fd := dtoType.Field(i)
		key := fd.Tag.Get("form")
		val := dtoVal.FieldByName(fd.Name)
	
		if !val.IsZero() && !fd.Anonymous{
			db = db.Where(fmt.Sprintf("%s = ?",key) ,val.Interface())
		}
	}
	return db
}

type PageDto struct {
	PageSize   int `form:"pageSize" json:"pageSize"  binding:"required"`
	Page int `form:"page" json:"page" binding:"required"`
}

func (page *PageDto) ResolvePage(db *gorm.DB)  *gorm.DB{
	return  db.Limit(page.PageSize).Offset(page.PageSize * (page.Page - 1))
}



