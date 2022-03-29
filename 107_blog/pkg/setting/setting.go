package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	/*
		Duration以int64纳秒计数表示两个瞬间之间经过的时间。
		表示将最大可代表的持续时间限制为大约290年。
	*/
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	/*
		从INI数据源加载和解析。参数可以是文件名和字符串
		类型的混合，也可以是Jbyte格式的原始数据。
		如果列表中包含不存在的文件，则返回错误。
	*/
	Cfg, err = ini.Load("107_blog/conf/app.ini")
	if err != nil {
		/*
			Fatalf相当于Printf()后面跟着对os.Exit(1)的调用。
		*/
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()

}
func LoadBase() {
	//Section假设指定的Section存在，当不存在时返回一个零值。
	//Key假设指定的Key存在于section中，当不存在时返回零值。
	//如果key值为空，MustString返回默认值。
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	//GetSection根据给定的名称返回section。
	sec, err := Cfg.GetSection("server")
	if err != nil {
		//Fatalf相当于Printf()后面跟着对os.Exit(1)的调用。
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	//返回key对应的值，不存在返回默认值
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	//获得app部分的数据
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
