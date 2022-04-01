// package setting

// //1读取配置文件的内容
// import (
// 	"log"
// 	"time"

// 	"gopkg.in/ini.v1"
// )

// var (
// 	Cfg *ini.File

// 	RunMode string

// 	HTTPPort int
// 	/*
// 		Duration以int64纳秒计数表示两个瞬间之间经过的时间。
// 		表示将最大可代表的持续时间限制为大约290年。
// 	*/
// 	ReadTimeout  time.Duration
// 	WriteTimeout time.Duration

// 	PageSize  int
// 	JwtSecret string
// )

// func init() {
// 	var err error
// 	/*
// 		从INI数据源加载和解析。参数可以是文件名和字符串
// 		类型的混合，也可以是Jbyte格式的原始数据。
// 		如果列表中包含不存在的文件，则返回错误。
// 	*/
// 	Cfg, err = ini.Load("conf/app.ini")
// 	if err != nil {
// 		/*
// 			Fatalf相当于Printf()后面跟着对os.Exit(1)的调用。
// 		*/
// 		log.Fatalf("Fail to get section 'server': %v", err)
// 	}
// 	LoadBase()
// 	LoadServer()
// 	LoadApp()

// }
// func LoadBase() {
// 	//Section假设指定的Section存在，当不存在时返回一个零值。
// 	//Key假设指定的Key存在于section中，当不存在时返回零值。
// 	//如果key值为空，MustString返回默认值。
// 	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
// }
// func LoadServer() {
// 	//GetSection根据给定的名称返回section。
// 	sec, err := Cfg.GetSection("server")
// 	if err != nil {
// 		//Fatalf相当于Printf()后面跟着对os.Exit(1)的调用。
// 		log.Fatalf("Fail to get section 'server': %v", err)
// 	}
// 	//返回key对应的值，不存在返回默认值
// 	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
// 	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
// 	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
// }

// func LoadApp() {
// 	//获得app部分的数据
// 	sec, err := Cfg.GetSection("app")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'app': %v", err)
// 	}

// 	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
// 	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
// }

package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

//加载配置文件
func Setup() {
	//加载文件
	Cfg, err := ini.Load("conf/app.ini")
	//如果加载出错
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	//获取对应的配置，放入对应的结构体中
	err = Cfg.Section("app").MapTo(AppSetting)
	//如果加载出错
	if err != nil {
		//打印日志并退出
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	//设置图像最大值
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	//获取server部分的值，使用mapto映射到结构体中
	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}
	//设置最大读取和写入时间
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	//获取db部分的值，使用mapto映射到结构体中
	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		//有错误打印错误并退出
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}
