package main

import (
	"GolangLearning/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"os"
)

func main() {
	// 加载配置
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()

	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册路由
	routers := routers.Register()

	// 加载模板文件
	//r.LoadHTMLGlob("templates/**/*")

	// 加载静态文件
	//r.Static("/static", "static")

	http_port := cfg.Section("").Key("http_port").String()

	routers.Run(http_port)
}
