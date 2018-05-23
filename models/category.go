package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
	State int `json:"state"`
}
//func ExistTagByName(name string) bool {
//	var category Category
//	db.Select("id").Where("name = ?  ", name, 0).First(&category)
//	if category.ID > 0 {
//		return true
//	}
//
//	return false
//}
func GetCategories(pageNum int, pageSize int, maps interface{}) (categories []Category) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&categories)
	return
}
func AddCategory(name string, state int) bool {
	db.Create(&Category{
		Name:      name,
		State:     state,

	})
	return true
}

func DeleteCategory(id int) bool {
	db.Where("id = ?", id).Delete(&Category{})

	return true
}