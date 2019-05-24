package Route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webchat/Controller/Home"
)

/**
	路由核心
 */

/**
	初始化路由
 */
func InitRoute() *gin.Engine  {

	route := gin.Default()


	route.LoadHTMLGlob("Views/**/*") //模板页面

	route.StaticFS("/static",http.Dir("Static/"))//静态资源

	route.GET("/index",Home.ChatIndex)

	//设置路由分组v1
	v1 := route.Group("/v1")

	{
		v1.GET("/login", func(context *gin.Context) {
			
		})

		v1.GET("/json/getList",Home.UseList)

		v1.GET("/json/getMembers",Home.UserMembers)
	}

	
	
	v2 := route.Group("/v2")
	{
		v2.GET("/login", func(context *gin.Context) {
			
		})
	}


	return route
}
