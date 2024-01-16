package common

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Dto struct {
}

func ResolveSearchQuery (ins interface{}) func(*gorm.DB)*gorm.DB{
	return func(db *gorm.DB)*gorm.DB{
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
}

//分页
type PageDto struct {
	PageSize   int `form:"pageSize" json:"pageSize"  binding:"required"`
	Page int `form:"page" json:"page" binding:"required"`
}

func (dto *PageDto)Paginate(db *gorm.DB)  *gorm.DB {
    page  := dto.Page
    if page <= 0 {
      page = 1
    }

    pageSize := dto.PageSize
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
}




