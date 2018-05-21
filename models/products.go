package models

type Product struct {
	Model
	Name  string
	Total int
	Left  int
	State int
}

func GetProducts(pageNum int, pageSize int, maps interface{}) (prouducts []Product) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&prouducts)
	return
}

func GetProductTotal(maps interface{}) (count int) {
	db.Model(&Product{}).Where(maps).Count(&count)

	return
}
