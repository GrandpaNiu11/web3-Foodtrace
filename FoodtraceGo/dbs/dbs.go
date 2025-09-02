package dbs

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func init() {
	fmt.Println("开始连接数据库")

	wd, err := os.Getwd()
	cfg, err := ini.Load(wd + "/my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开启SQL日志
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接数据库成功")
}
func loadConfig() (*ini.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取工作目录失败: %v", err)
	}
	configPath := wd + "/my.ini" // 基于工作目录
	cfg, err := ini.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("加载配置文件失败: %v", err)
	}
	return cfg, nil
}
