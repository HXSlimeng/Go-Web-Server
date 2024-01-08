package common

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Dto struct {
}

func (dto *Dto) ResolveSearchQuery(db *gorm.DB) *gorm.DB {
	st := reflect.TypeOf(dto)
	for i := 0; i < st.NumField(); i++ {
		str := st.Field(i).Tag.Get("search")
		fmt.Println(str)
	}
	
	return db 
}