package main
import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "I am guiltyfet",
	})
}
func main(){
	r := gin.Default() //返回默认的路由引擎

	//指定用户使用get请求访问
	r.GET("/hello", sayHello)

	r.Run()
}
