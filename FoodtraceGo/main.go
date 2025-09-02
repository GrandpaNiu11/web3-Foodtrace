package main

import (
	"FoodtraceGo/intercept"
	"FoodtraceGo/router"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(intercept.ErrorHandler())
	// 自定义 CORS 配置
	// 允许多个来源
	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // localhost
			"http://127.0.0.1:3000", // 127.0.0.1
			"http://0.0.0.0:3000",   // 0.0.0.0
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))
	router.BcosRoutersInit(r)
	// 示例2: 生成私钥和地址

	fmt.Println(r.Run(":8080"))

}
