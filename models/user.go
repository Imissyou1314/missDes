package models

// 用户信息表
type User struct {
	Id int64
	// 用户Id
	UserId int64
	//用户销售Id
	UserSaleLogId string
	//用户名字
	Username string
	//用户密码
	Password string
	//用户当前状态
	Status bool
}

// 获取指定用户
func GetUser(user User) {
	GetInstance(user)
}

// 获取所有用户
func GetAllUsers() []User {
	var users []User
	GetQueryByTableName("user").All(users)
	Logger.Println("<=== Users: ===>", users)
	return users
}

// 更新用户
func UpdateUser(user User) {
	Logger.Println("<=== user ===>", user)
	UpdateInstace(user)
}

// 删除指定用户
func DeleteUser(user User) {
	Logger.Println("<=== user ===>", user)
	DeleteInstace(user)
}

// 添加指定用户
func InsertUser(user User) {
	Logger.Println("<=== user ===>", user)
	InsertInstance(user)
}
