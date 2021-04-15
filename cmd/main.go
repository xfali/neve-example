// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package main

import (
	"github.com/xfali/neve-core"
	"github.com/xfali/neve-core/processor"
	"github.com/xfali/neve-database/gobatiseve"
	_ "github.com/xfali/neve-example/docs/swagdoc"
	"github.com/xfali/neve-example/internal/pkg/cache"
	"github.com/xfali/neve-example/pkg/swagger"
	"github.com/xfali/neve-logger/xlogneve"
	"github.com/xfali/neve-utils/neverror"
	"github.com/xfali/neve-web/gineve"
	"github.com/xfali/xlog"
)

var ConfigFile string

func init() {
	if neve.ResourceRoot == "" {
		neve.SetResourceRoot("assets")
	}
	if ConfigFile == "" {
		ConfigFile = "config-example.yaml"
	}
}

// Xiongfa Li

// @title Neve Example
// @version v1.0.0
// @description Neve Example

// @contact.name Xiongfa.Li
// @contact.email lxf1222@163.com
func main() {
	confFile := neve.GetResource(ConfigFile)
	xlog.Infoln("Config file: ", confFile)
	app := neve.NewFileConfigApplication(confFile)
	// 配置日志处理器，读取配置并初始化日志系统
	neverror.PanicError(app.RegisterBean(xlogneve.NewLoggerProcessor()))
	// 配置值注入处理器，添加值注入功能。此处配置标识注入的tag为value
	neverror.PanicError(app.RegisterBean(processor.NewValueProcessor(processor.OptSetValueTag("", "value"))))
	// 配置web服务处理器，添加web服务路由注册功能
	neverror.PanicError(app.RegisterBean(gineve.NewProcessor()))
	// 配置数据库处理器，读取配置生成datasource
	neverror.PanicError(app.RegisterBean(gobatiseve.NewProcessor()))
	// 配置swagger
	neverror.PanicError(app.RegisterBean(swagger.NewHandler()))

	// 自定义的cache service
	neverror.PanicError(app.RegisterBean(cache.NewService()))
	// if use mysql:
	//neverror.PanicError(app.RegisterBean(mysql.NewService()))
	// 自定义的cache handler
	neverror.PanicError(app.RegisterBean(cache.NewHandler()))

	app.Run()
}
