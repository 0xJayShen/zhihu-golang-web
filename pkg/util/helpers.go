package util

import (
	"strings"
	"strconv"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
)

func Merge_name_helper(prefix string, name interface{}) string {
	switch v := name.(type) {
	case int:
		var tmp int
		tmp = v
		merge_name := strings.Join([]string{prefix, strconv.Itoa(tmp)}, "_")
		return merge_name
	case string:
		var tmp string
		tmp = v
		merge_name := strings.Join([]string{prefix, tmp}, "_")
		return merge_name
	default:

		return "error"
	}
}


func ToMap(in2 interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	inrec, _ := json.Marshal(in2)
	json.Unmarshal(inrec, &m)
	return m
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
