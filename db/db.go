package db

import (
	"fmt"
	"github.com/Sion-L/onePiece/model"
	ldapv3 "github.com/go-ldap/ldap/v3"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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

func NewClientDB() (*gorm.DB, error) {
	var config DatabaseConfig
	err := configor.Load(&config, "./config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %s", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err)
	}
	db.AutoMigrate(&model.User{})
	return db, nil
}

func NewClientLdap() (*ldapv3.Conn, error) {
	var config LdapConfig
	err := configor.Load(&config, "./config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %s", err)
	}
	conn, err := ldapv3.Dial("tcp", fmt.Sprintf("%s:%v", config.Ldap.Host, config.Ldap.Port))
	if err != nil {
		return nil, err
	}
	err = conn.Bind(config.Ldap.BindDn, config.Ldap.BindPw)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
