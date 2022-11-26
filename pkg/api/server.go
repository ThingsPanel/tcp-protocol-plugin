package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type apiServer struct {
	r    *gin.Engine
	addr string
}

func NewCustomServer(addr string) *apiServer {
	e := gin.Default()
	e.POST("/api/form/config", CustomGetFormConfig)
	e.POST("/api/device/config/update", CustomDeviceConfigUpdate)
	e.POST("/api/device/config/add", CustomDeviceConfigCreate)
	e.POST("/api/device/config/delete", CustomDeviceConfigDelete)
	return &apiServer{
		r:    e,
		addr: addr,
	}
}

func NewSelfApiServer(addr string) *apiServer {
	e := gin.Default()
	e.POST("/api/form/config", SelfGetFormConfig)
	e.POST("/api/device/config/update", SelfDeviceConfigUpdate)
	e.POST("/api/device/config/add", SelfDeviceConfigCreate)
	e.POST("/api/device/config/delete", SelfDeviceConfigDelete)
	return &apiServer{
		r:    e,
		addr: addr,
	}
}

func (s *apiServer) Start() {
	log.Info("http server started at:" + s.addr)
	s.r.Run(s.addr)
}
