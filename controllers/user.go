package controllers

import (
	// "encoding/json"
	"missDes/models"

	"missDes/utils"
	// "strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func getInputs(t beego.Controller) map[string]string {
	return t.Ctx.Input.Params()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) CreateUser() {
	inputs := u.Ctx.Input
	var user models.User
	user.Username = inputs.Param("userName")
	user.Password = utils.GetMD5Hash(inputs.Param("password"))
	user.UserSaleLogId = utils.GetMD5Hash(inputs.Param("userName"))
	models.InsertUser(user)
}

func (u *UserController) GetAllUser() {
	var users []modeks.User
	users = models.GetAllUsers()
	u.Data["json"], _ = utils.ToJSONStr(users)
}

func (u *UserController) UpUser() {
	inputs := u.Ctx.Input
	var user models.User
	user.Username = inputs.Param("userName")
	user.Password = utils.GetMD5Hash(inputs.Param("password"))
	user.UserSaleLogId = utils.GetMD5Hash(inputs.Param("userName"))
	models.UpdateUser(user)
}
