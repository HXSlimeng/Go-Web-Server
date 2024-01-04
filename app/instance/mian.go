package instance

import (
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	Instance *gin.Engine 
}

func (ins *WebServer) Create(){
	ins.Instance = gin.Default()
}
func (ins *WebServer) Start(path string) error  {
	return ins.Instance.Run(path)
}

func (ins *WebServer) SetupRouter(test func()){
	
}


