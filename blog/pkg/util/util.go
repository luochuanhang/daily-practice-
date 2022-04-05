package util

import "lianxi/blog/pkg/setting"

// Setup 设置JWT秘钥
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
