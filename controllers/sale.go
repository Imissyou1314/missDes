package controllers

import (
	"missDes/models"

	"strconv"

	"time"

	"github.com/astaxie/beego"
)

// the SaleController extends Beego Controller
type SaleController struct {
	beego.Controller
}

func (s *SaleController) CreateSaleLog() {
	inputs := s.Ctx.Input
	var saleLog models.SaleLog

	saleLog.SaleNumber, _ = strconv.Atoi(inputs.Param("saleNumber"))
	saleLog.SaleDesc = inputs.Param("saleDesc")
	saleLog.SaleTime = time.Now().String()
}
