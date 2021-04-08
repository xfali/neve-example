// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	Enable bool   `value:"swagger.enable"`
	Path   string `value:"swagger.path"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HttpRoutes(engine gin.IRouter) {
	if h.Enable {
		if h.Path == "" {
			h.Path = "/swagger"
		}
		engine.GET(h.Path+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
