package gormx

import (
	"gorm.io/gorm"
	"reflect"
)

//分页功能
func Paginate(page interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := reflect.ValueOf(page)
		pageNum := page.FieldByName("PageNum").Interface().(int) - 1
		pageSize := page.FieldByName("PageSize").Interface().(int)
		offset := pageNum * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
