package models

import (
	"zhihu-golang-web/cache"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"zhihu-golang-web/pkg/logging"
	"time"
	"zhihu-golang-web/pkg/msg"
)

type Product struct {
	gorm.Model
	CreateBy   string   `json:"create_by"`
	Name       string   `json:"name"`
	Total      int      `json:"total"`
	Left       int      `json:"left"`
	State      int      `json:"state"`
	Des        string   `json:"des"` //列名是 des
	CategoryId int      `json:"category_id"`
	Category   Category `json:"category"`
}

func GetProducts(pageNum int, pageSize int, maps interface{}) (products []Product) {
	db.Preload("Category").Where(maps).Offset(pageNum).Limit(pageSize).Find(&products)
	return
}
func GetProductTotal(maps interface{}) (count int) {
	conn := cache.RedisPool.Get()
	defer conn.Close()
	is_key_exit, err := redis.Bool(conn.Do("EXISTS", "product_total"))
	if err != nil {
		logging.Error("缓存查商品总量错误", time.Now().Format(msg.TIME_FORMAT))
	}
	if is_key_exit {
		count, err = redis.Int(conn.Do("GET", "product_total"))
		if err != nil {
			logging.Error("缓存总量转换int 类型出错")
		}
		return
	} else {
		db.Model(&Product{}).Where(maps).Count(&count)
		_, e := conn.Do("SET", "product_total", count, "EX", "100000")
		if e != nil {
			logging.Error("缓存存入商品总量出错")
		}

	}
	return
}
func IsProductCacheExit(id int) bool {
	conn := cache.RedisPool.Get()
	defer conn.Close()
	is_key_exit, err := redis.Bool(conn.Do("EXISTS", id))
	if err != nil {
		logging.Error("缓存中查询商品出错 %v", id, err, time.Now().Format(msg.TIME_FORMAT))
	}
	if is_key_exit {
		return true
	} else {
		return false
	}
}
func IsProductDBExit(id int) bool {
	var product Product
	db.Select("id").Where("id = ? ", id, ).First(&product)
	if product.ID > 0 {
		return true
	} else {
		return false
	}
}

func GetProduct(id int) interface{} {
	exit := IsProductCacheExit(id)
	conn := cache.RedisPool.Get()
	defer conn.Close()
	if exit {
		var info_map map[string]interface{}
		valueGet, err := redis.Bytes(conn.Do("GET", id))

		if err != nil {
			logging.Error("缓存中取 %v 出错", id, err, time.Now().Format(msg.TIME_FORMAT))
		}

		err = json.Unmarshal(valueGet, &info_map)

		if err != nil {
			logging.Error("反序列化出错 %v", id, err, time.Now().Format(msg.TIME_FORMAT))
		}
		return info_map

	} else {
		var product Product
		err := db.Where("id = ? ", id).First(&product)
		key_name := strings.Join([]string{"shop_product", strconv.Itoa(id)}, "_")
		if err.Error == gorm.ErrRecordNotFound {
			fmt.Println("找不到商品")
			valu_map := map[string]int{"id": id, "state": -1}
			value_json, err := json.Marshal(valu_map)
			_, err = conn.Do("SET", key_name, value_json)
			if err != nil {
				fmt.Println(err, 333333)
			}
			return valu_map
		}
		product_json, _ := json.Marshal(product)

		_, e := conn.Do("SET", key_name, product_json, "EX", "100000")
		if e != nil {
			logging.Error("存缓存出错 %v", id, time.Now().Format(msg.TIME_FORMAT))
		}
		return product
	}
}
func DeleteProduct(id int) bool {
	key_name := strings.Join([]string{"shop_product", strconv.Itoa(id)}, "_")
	conn := cache.RedisPool.Get()
	exit := IsProductCacheExit(id)
	defer conn.Close()

	if exit {
		conn.Do("DEL", key_name)
		db.Where("id = ?", id).Delete(Product{})
	} else {
		db.Where("id = ?", id).Delete(Product{})
	}
	return true
}

func BuyProduct(id int) (product Product) {
	conn := cache.RedisPool.Get()
	defer conn.Close()

	conn.Do("Watch")
	exit := IsProductCacheExit(id)
	if exit {
		conn.Do("DEL", id)
	} else {

	}
	db.Exec("update blog_product set left=left - 1 where id = 1")
	return
}
