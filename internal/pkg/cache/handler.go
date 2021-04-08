// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package cache

import (
	"github.com/gin-gonic/gin"
	"github.com/xfali/neve-web/gineve/midware"
	"github.com/xfali/xlog"
	"net/http"
)

type Handler struct {
	log  xlog.Logger
	Svc  Service            `inject:""`
	HLog midware.HttpLogger `inject:""`
}

func NewHandler() *Handler {
	return &Handler{
		log: xlog.GetLogger(),
	}
}

func (h *Handler) HttpRoutes(engine gin.IRouter) {
	engine.GET("/keys/:key", h.HLog.LogHttp(), h.get)
	engine.PUT("/keys/:key", h.HLog.LogHttp(), h.set)
}

func (h *Handler) get(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	v := h.Svc.Get(key)
	if v == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.log.Infof("Get key: %s , value: %s\n", key, v)
	ctx.Writer.Write([]byte(v))
}

func (h *Handler) set(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	v, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	h.log.Infof("Set key: %s , value: %s\n", key, string(v))
	h.Svc.Set(key, string(v))
}
