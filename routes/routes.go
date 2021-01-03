/*
 * @Author: Jecosine
 * @Date: 2021-01-03 07:53:33
 * @LastEditTime: 2021-01-04 06:34:28
 * @LastEditors: Jecosine
 * @Description: router definition
 */
package routes

import (
	"github.com/Jecosine/obm-back-end/pkg/setting"
	v1 "github.com/Jecosine/obm-back-end/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	// use logger
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/user/get/:id", v1.GetUser)
		apiV1.POST("/user/save/:id", v1.EditUser)
		apiV1.POST("/user/del/:id", v1.DeleteUser)
		apiV1.POST("/user/add", v1.AddUser)
	}
	return r
}
