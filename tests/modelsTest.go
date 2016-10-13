package test

import (
	"testing"

	"missDes/models"

	"github.com/astaxie/beego/logs"
)

var Logger = logs.GetLogger()

func init() {
	Logger.Println("<==== this is Test Models ====>")
}

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
