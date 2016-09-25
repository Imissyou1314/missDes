package models

type SaleLog struct {
	Id int `orm:"auto"`
	//  销售数量
	SaleNumber int
	//  销售时间
	SaleTime string `orm:"size(100)"`
	//  销售说明
	SaleDesc string `orm:"size(100)"`
}
