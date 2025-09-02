package router

import (
	"FoodtraceGo/controller"
	"github.com/gin-gonic/gin"
)

func BcosRoutersInit(r *gin.Engine) {
	r.POST("/addLogisticsProvider", controller.FoodController{}.AddLogisticsProvider)
	r.POST("/registerFood", controller.FoodController{}.AddRegisterFood)
	r.POST("/updateLogistice", controller.FoodController{}.AddUpdateLogistice)
	r.GET("/logisticsRecordedCount", controller.FoodController{}.LogisticsRecordedCount)
	r.GET("/foodItemCount", controller.FoodController{}.FoodItemCount)
	r.GET("/foodItems", controller.FoodController{}.FoodItems)
	r.GET("/logisticsProviders", controller.FoodController{}.LogisticsProviders)
	r.GET("/getFoodInfo", controller.FoodController{}.GetFoodInfo)
	r.GET("/selectFoodItems", controller.FoodController{}.SelectFoodItems)
	r.POST("/userRegister", controller.AdminController{}.UserRegister)
	r.POST("/userLogin", controller.AdminController{}.UserLogin)
}
