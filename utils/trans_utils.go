package utils

import (
	"strconv"
	"encoding/json"
)

func Bool2String(v bool) string {
	return strconv.FormatBool(v)
}

func Int2String(i int) string {
	return strconv.Itoa(i)
	//return strconv.FormatInt(i,32)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i,64)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f,'E',-1,32)
}

func String2Int(v string) (int,error) {
	return strconv.Atoi(v)
}

func String2Int64(v string) (int64,error) {
	return strconv.ParseInt(v,10,64)
}

func String2Float(v string) (float64,error){
	return strconv.ParseFloat(v,64)
}

//v是json字符串，i是struct地址或者map地址
func JsonString2StructOrMap(v string,i interface{}) error {
	return json.Unmarshal([]byte(v),i)
}

func StructOrMap2JsonString(i interface{}) (string,error) {
	data,err := json.Marshal(i)
	return string(data),err
}

