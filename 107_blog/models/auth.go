package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//检查数据库有没有这个账号密码
func CheckAuth(username, password string) bool {
	var auth Auth
	//查询对应的账号密码对应的id
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	//如果有数据id>0
	return auth.ID > 0
}
