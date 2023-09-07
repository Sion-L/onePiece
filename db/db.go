package db

import (
	"fmt"
	"github.com/Sion-L/onePiece/model"
	ldapv3 "github.com/go-ldap/ldap/v3"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Conn *gorm.DB
var LdapConn *ldapv3.Conn

type DatabaseConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

type LdapConfig struct {
	Ldap struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		BaseDn string `yaml:"baseDn"`
		BindDn string `yaml:"bindDn"`
		BindPw string `yaml:"bindPw"`
	} `yaml:"ldap"`
}

func InitMySQLDB() {
	var config DatabaseConfig
	err := configor.Load(&config, "./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("error loading config: %s", err))
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	Conn, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %s", err))
	}
	sqlDB := Conn.DB()
	sqlDB.SetMaxIdleConns(10)                  // 连接池最大闲置连接
	sqlDB.SetMaxOpenConns(100)                 // 设置最大连接数
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置连接的最大可复用时间
	Conn.AutoMigrate(&model.User{})
}

func InitLdap() {
	var config LdapConfig
	err := configor.Load(&config, "./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("error loading config: %s", err))
	}
	LdapConn, err = ldapv3.Dial("tcp", fmt.Sprintf("%s:%v", config.Ldap.Host, config.Ldap.Port))
	if err != nil {
		panic(fmt.Errorf("error connecting to ldap: %s", err))
	}
	err = LdapConn.Bind(config.Ldap.BindDn, config.Ldap.BindPw)
	if err != nil {
		panic(fmt.Errorf("error binding to ldap: %s", err))
	}
}
