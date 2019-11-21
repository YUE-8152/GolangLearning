package db

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	var err error
	var cfg *ini.File
	var maxIdleConns int
	var maxOpenConns int

	// load配置
	cfg, err = ini.Load("conf/database.ini", "conf/app.ini")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()
	// 主机
	host := cfg.Section(mode).Key("mysql.host").String()
	// 端口
	port := cfg.Section(mode).Key("mysql.port").String()
	// 用户名
	username := cfg.Section(mode).Key("mysql.username").String()
	// 密码
	password := cfg.Section(mode).Key("mysql.password").String()
	// 数据库名称
	dbname := cfg.Section(mode).Key("mysql.dbname").String()
	// 最大空闲连接数
	maxIdleConns, err = cfg.Section(mode).Key("mysql.max_idle_conns").Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 最大打开的连接数
	maxOpenConns, err = cfg.Section(mode).Key("mysql.max_open_conns").Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	//是否默认数据库表加"s"
	singularTable, err := cfg.Section(mode).Key("mysql.singular_table").Bool()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	//是否打印sql语句
	logMode, err := cfg.Section(mode).Key("mysql.log_mode").Bool()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Fail to open mysql: %v", err)
		os.Exit(1)
	}

	db.SingularTable(singularTable)
	db.LogMode(logMode)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetConnMaxLifetime(time.Hour)

	//db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//	fmt.Println("数据库连接失败！")
	//}
	//db.SingularTable(true)
	//db.LogMode(true)
	//fmt.Println("数据库连接成功！")
	DB = db
}
