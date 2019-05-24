package main

import (
	"fmt"
	"webchat/Core"
	"webchat/Route"
)

/**
	入口函数
 */
func main()  {


	route := Route.InitRoute()

	addr := Core.NewConfig()

	fmt.Print(addr.Port)

	route.Run(":8088")

}