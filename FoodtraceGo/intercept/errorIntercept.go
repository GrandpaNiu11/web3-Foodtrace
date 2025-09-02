package intercept

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 统一错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先执行后续的中间件和处理函数

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// 可以根据错误类型进行不同的处理
			log.Printf("请求失败: %v", err)

			// 统一返回错误格式
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}
	}
}
