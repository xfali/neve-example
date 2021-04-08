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

	Version string `value:"version"`
}

func NewHandler() *Handler {
	return &Handler{
		log: xlog.GetLogger(),
	}
}

func (h *Handler) HttpRoutes(engine gin.IRouter) {
	engine.GET("/version", h.HLog.LogHttp(), h.version)
	engine.GET("/keys/:key", h.HLog.LogHttp(), h.get)
	engine.PUT("/keys/:key", h.HLog.LogHttp(), h.set)
	engine.DELETE("/keys/:key", h.HLog.LogHttp(), h.delete)
}

func (h *Handler) version(ctx *gin.Context) {
	ctx.Writer.Write([]byte(h.Version))
}

// @Summary Get Value
// @Description Get Value
// @Tags cache
// @Param key path string true "key"
// @Success 200 {array} string
// @Failure 400 {array} string
// @Failure 404 {array} string
// @Router /keys/{key} [get]
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
	ctx.Writer.WriteString(v)
}

// @Summary Add Key and Value
// @Description Add Key and Value
// @Tags cache
// @Success 200 {array} string
// @Failure 400 {array} string
// @Router /keys/{key} [put]
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

// @Summary Delete Key and Value
// @Description Delete Key and Value
// @Tags cache
// @Param key path string true "key"
// @Success 200 {array} string
// @Failure 400 {array} string
// @Failure 404 {array} string
// @Router /keys/{key} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ok := h.Svc.Delete(key)
	if !ok {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.log.Infof("Delete key: %s\n", key)
}
