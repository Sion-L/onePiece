package db

import (
	"fmt"
	"github.com/Sion-L/onePiece/model"
	ldapv3 "github.com/go-ldap/ldap/v3"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func NewClientDB() {
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
	Conn.AutoMigrate(&model.User{})
}

func NewClientLdap() {
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
