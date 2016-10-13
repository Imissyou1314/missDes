package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"

	_ "github.com/go-sql-driver/mysql"
)

var Logger = logs.GetLogger("models")

func init() {
	// 设置控制台打印日子
	logs.SetLogger("console")
	//  控制日志输出文件为models.logs
	logs.SetLogger(logs.AdapterFile,
		`{"filename":"models.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//  输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	//  设置日志缓存的大小
	logs.Async(1e3)
	//  注册数据库连接类型
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//  注册数据库连接配置
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/miss?charset=utf8")

	orm.RegisterModel(new(Address), new(SaleLog), new(User))
}

//  创建Orm 并配置使用的数据库
// return: orm.Ormer
func getOrm() orm.Ormer {
	o := orm.NewOrm()
	// 使用默认数据库
	o.Using("default")
	return o
}

//	checkError:检查是否出错
//	@param:err 错误谢谢
//	@return: if have err return true or return false that the err is nil
func checkError(err error) bool {
	if err != nil {
		logs.Error("miss", err)
		return true
	} else {
		return false
	}
}

// GetInstance 通过Id获取数据
// instance: 查询的对象
func GetInstance(instance interface{}) {
	o := getOrm()
	o.Read(&instance)
}

//  InsertInstance插入数据
//  instance:插入的对象
func InsertInstance(instance interface{}) {
	o := getOrm()
	o.Insert(instance)
}

//	DeleteInstace 删除数据
//  instance:删除数据库中的该对象
func DeleteInstace(instance interface{}) {
	o := getOrm()
	o.Delete(instance)
}

//	UpdateInstace:更新对应得对象
//  @param: instance：对应的对象
//  @return：int64  放回更改得到行数
func UpdateInstace(instance interface{}) int64 {
	o := getOrm()
	changeNumber, err := o.Update(instance)
	checkError(err)
	return changeNumber
}

//	GetQueryByTableName
// 获取操作表
func GetQueryByTableName(tableName string) orm.QuerySeter {
	return getOrm().QueryTable(tableName)
}
