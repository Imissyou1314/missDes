package models

//	销售记录表
type SaleLog struct {
	Id int64 `orm:"auto"`
	//  销售数量
	SaleNumber int
	//  销售时间
	SaleTime string `orm:"size(100)"`
	//  销售说明
	SaleDesc string `orm:"size(100)"`
}

//	获取指定销售记录信息
func GetSale(sale SaleLog) SaleLog {
	GetInstance(sale)
	return sale
}

//	获取所有销售记录信息
func GetAllSale() []*SaleLog {
	var saleLogs []*SaleLog
	GetQueryByTableName("sale_log").All(&saleLogs)
	return saleLogs
}

//	添加一条销售记录信息
func InsertSale(sale SaleLog) {
	InsertInstance(sale)
}

//	删除指定的销售信息
func DeleteSale(sale SaleLog) {
	DeleteInstace(sale)
}
