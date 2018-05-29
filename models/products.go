package models

import (
	"github.com/garyburd/redigo/redis"
	"github.com/qq976739120/zhihu-golang-web/cache"
	"fmt"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/qq976739120/zhihu-golang-web/pkg/logging"
	"time"
	"github.com/qq976739120/zhihu-golang-web/pkg/msg"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
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

	key_name := util.Merge_name_helper("shop_product", id)

	is_key_exit, err := redis.Bool(conn.Do("EXISTS", key_name))
	if err != nil {
		logging.Error("缓存中查询商品出错", id, err, time.Now().Format(msg.TIME_FORMAT))
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
	key_name := util.Merge_name_helper("shop_product", id)

	if exit {
		var res map[string]interface{}
		var category_json map[string]interface{}
		res = make(map[string]interface{})
		res_bytes, _ := redis.StringMap(conn.Do("HGETALL", key_name))
		for k, v := range res_bytes {
			res[k] = v
		}
		json.Unmarshal([]byte(res_bytes["category"]), &category_json)
		res["category"] = category_json
		return res
	} else {
		var product Product
		err := db.Preload("Category").Where("id = ? ", id).First(&product)
		if err.Error == gorm.ErrRecordNotFound {
			valu_map := map[string]int{"id": id, "state": -1}
			value_json, err := json.Marshal(valu_map)
			_, err = conn.Do("SET", key_name, value_json)
			if err != nil {
				fmt.Println(err, 333333)
			}
			return valu_map
		}
		category_json, _ := json.Marshal(util.ToMap(product.Category))
		m := map[string]interface{}{
			"id":       product.ID,
			"name":     product.Name,
			"total":    product.Total,
			"left":     product.Left,
			"state":    product.State,
			"des":      product.Des,
			"category": category_json,
		}
		_, er := conn.Do("HMSET", redis.Args{}.Add(key_name).AddFlat(m)...)
		fmt.Println(er)
		return m
	}
}
func DeleteProduct(id int) bool {
	key_name := util.Merge_name_helper("shop_product", id)
	conn := cache.RedisPool.Get()
	defer conn.Close()
	db.Where("id = ?", id).Delete(Product{})
	conn.Do("DEL", key_name)

	return true
}

func BuyProduct(id int) bool {
	conn := cache.RedisPool.Get()
	defer conn.Close()
	exit := IsProductCacheExit(id)
	key_name := util.Merge_name_helper("shop_product", id)
	if exit {
		conn.Do("hincrby", key_name, "left", -1)
	} else {
		exit = IsProductDBExit(id)
		if exit {
			db.Exec("update shop_student set left=left-1")
		} else {
			return false
		}
	}
	return true
}
func SecondKill(user_id int)bool {
	//总量
	var len_list int
	all_num  := 100
	conn := cache.RedisPool.Get()
	defer conn.Close()
	len_reply,err := conn.Do("LLEN","second_kill")
	switch v := len_reply.(type) {
	case int:
		len_list = v
	}
	if len_list >= all_num{
		return false
	}else{
		_, err = conn.Do("LPUSH", "second_kill", user_id)
		return true
	}
	if err != nil{
		logging.Error("秒杀出错", err)
	}
	return false
}
