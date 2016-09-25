package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

var logger = logs.GetLogger()

//  构造函数，初始化Logger
func init() {
	fmt.Println("<==== json 构造 函数 ======>")
	logger.Println("配置json Utils 初始化....")
}

//  ToJSON 将字符串转成json
// 对象转Json字符串
// jsonData
func ToJSON(jsonData interface{}) (b []byte, err error) {
	logger.Println("to json data is:", jsonData)
	if jsonData == nil {
		logger.Println("json data is nil")
	} else {
		b, err := json.Marshal(jsonData)
		return b, err
	}
	return nil, nil
}

//  JSONToStruct 数据转换
//  json数据转成对象
//  jsonData
//  instance
func JSONToStruct(jsonData []byte, instance interface{}) error {
	logger.Println("json to Struct:", jsonData)
	if jsonData == nil {
		logger.Println("json_data is nil")
	} else {
		err := json.Unmarshal(jsonData, &instance)
		return err
	}
	return nil
}
