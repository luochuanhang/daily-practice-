package models

import "github.com/jinzhu/gorm"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//CheckAuth检查认证信息是否存在
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	//查询数据库中账号密码对应的id
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	//如果id大于0,返回true
	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
