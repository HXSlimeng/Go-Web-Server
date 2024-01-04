package sysdept

import "github.com/gin-gonic/gin"

// type Controller struct{}

func GetList(c *gin.Context) {
	sysDeptService := new(Service)
	list := sysDeptService.GetList()
	c.JSON(200,gin.H{"data":list})
}