package app

import (
	"github.com/astaxie/beego/validation"

	"lianxi/blog/pkg/logging"
)

// MarkErrors记录错误日志
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		//将错误记录到日志
		logging.Info(err.Key, err.Message)
	}

}
