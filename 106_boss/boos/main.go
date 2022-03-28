package main

import (
	"lianxi/106_boss/boos/pkg/config"
	"lianxi/106_boss/boos/pkg/router"
	"lianxi/106_boss/boos/pkg/storage"

	"github.com/sirupsen/logrus"
)

func main() {
	// 路由管理器
	var (
		engine = storage.NewStorage("default")
	)
	if err := config.Load("106_boss/boos/boos.json"); err != nil {
		logrus.WithError(err).Errorf("配置加载失败")
		panic("配置加载失败")
	}
	conf := config.DefaultConfig
	if conf.Mysql.String() != "" {
		if err := engine.Init(conf.Mysql.String()); err != nil {
			logrus.WithError(err).Errorf("初始化mysql失败")
		}
	}
	if conf.Redis.String() != "" {
		if err := storage.Init(conf.Redis.Addr, conf.Redis.Password, conf.Redis.DB); err != nil {
			logrus.WithError(err).Errorf("初始化redis失败")
		}
	}
	if err := router.Start(engine); err != nil {
		logrus.WithError(err).Errorf("http服务运行失败")
	}
}
