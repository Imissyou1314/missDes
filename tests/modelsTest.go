package test

import (
	"testing"

	"missDes/models"

	"github.com/astaxie/beego/logs"
)

//  记录测试日志的Logger
var Logger = logs.GetLogger()

func init() {
	Logger.Println("<==== this is Test Models ====>")
}

//  插入数据测试
func TestInsert(t *testing.T) {
	var address = new(models.Address)
	address.AddressName = "湛江"
	address.AddressNumber = 0
	address.AddressSaleNumber = 1
	models.InsertInstance(address)
}

func TestDelete(t *testing.T) {
	var user = new(models.User)
	user.Id = 1
	user.Password = "missyou"
	user.Username = "Imissyou"
	models.InsertInstance(user)
}
