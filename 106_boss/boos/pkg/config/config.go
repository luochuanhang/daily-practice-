package config

import (
	"encoding/json"
	"io/ioutil"

	"lianxi/106_boss/boos/pkg/storage"
)

type Config struct {
	Addr  string              `json:"addr"`
	Mysql storage.MysqlConfig `json:"mysql"`
	Redis storage.RedisConfig `json:"redis"`
}

var DefaultConfig Config

func Load(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var temp = &Config{}
	err = json.Unmarshal(bs, temp)
	if err != nil {
		return err
	}
	DefaultConfig = *temp
	return nil
}
