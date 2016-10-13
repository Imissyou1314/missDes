package models

//  Address models
//  记录地区的销售记录
type Address struct {
	Id int64 `orm:"auto"`
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
	//  销售介绍
	AddressSaleDes string
}

//	获取特定地址
func (a Address) GetAddress(address Address) {
	GetInstance(address)
}

//	获取所以得地址
func GetAllAddress() []*Address {
	var addresss []*Address
	GetQueryByTableName("address").All(addresss)
	if addresss != nil {
		return addresss
	} else {
		return nil
	}
}

//	删除指定地址
func DeleteAddress(address Address) {
	DeleteInstace(address)
}

//	插入指点地址
func InsertAddress(address Address) {
	InsertInstance(address)
}

//	更新指定地址
func UpdataAddress(address Address) {
	UpdateInstace(address)
}
