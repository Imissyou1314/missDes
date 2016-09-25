package models

//  Address models
//  记录地区的销售记录
type Address struct {
	Id int `orm:"auto"`
	//  地区的名字
	AddressName string `orm:"size(100)"`
	//  地区维度
	AddressLog float32
	//  地区经度
	AddressLan float32
	//  在该地区进行销售的次数
	AddressNumber int
	//  该地区销售的数量
	AddressSaleNumber int
	//  最后一次销售记录时间
	AddressSaleLateTime string
}

func GetAddress() *Address {
	return new(Address)
}
