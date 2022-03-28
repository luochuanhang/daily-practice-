package storage

import (
	"fmt"
	"sync"

	"lianxi/106_boss/boos/pkg/model"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlConfig stands for database connection configuration.
type MysqlConfig struct {
	User      string `json:"user,omitempty"`
	Password  string `json:"password,omitempty"`
	Host      string `json:"host,omitempty"`
	DbName    string `json:"db_name,omitempty"`
	Charset   string `json:"charset,omitempty"`
	ParseTime string `json:"parse_time,omitempty"`
	Loc       string `json:"loc,omitempty"`
}

func (m *MysqlConfig) String() string {
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		m.User, m.Password, m.Host, m.DbName, m.Charset, m.ParseTime, m.Loc)
}

type mysql struct {
	db *gorm.DB
}

var (
	defaultmysql = &mysql{}
	onceMysql    sync.Once
)

var _ Storage = &mysql{}

func (m *mysql) Name() string {
	return "mysql"
}

// Init init mysql DB connection.
func (m *mysql) Init(dns string) error {
	var err error
	onceMysql.Do(func() {
		if defaultmysql.db == nil {
			defaultmysql.db, err = gorm.Open(mysqldriver.Open(dns), &gorm.Config{})
		}
	})
	if err != nil {
		return fmt.Errorf("gorm.Open %w", err)
	}
	return defaultmysql.db.AutoMigrate(&model.User{})
}

// Get get default mysql connection.
func (m *mysql) Get() *gorm.DB {
	return defaultmysql.db
}

// Get get default mysql connection.
func (m *mysql) New(dns string) (*gorm.DB, error) {
	return gorm.Open(mysqldriver.Open(dns), &gorm.Config{})
}
