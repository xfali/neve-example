// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package main

import (
	"github.com/xfali/neve-core"
	"github.com/xfali/neve-core/processor"
	"github.com/xfali/neve-database/gobatiseve"
	"github.com/xfali/neve-example/internal/pkg/cache"
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

func main() {
	confFile := neve.GetResource(ConfigFile)
	xlog.Infoln("Config file: ", confFile)
	app := neve.NewFileConfigApplication(confFile)
	neverror.PanicError(app.RegisterBean(processor.NewValueProcessor(processor.OptSetValueTag("", "value"))))
	neverror.PanicError(app.RegisterBean(gineve.NewProcessor()))
	neverror.PanicError(app.RegisterBean(gobatiseve.NewProcessor()))

	neverror.PanicError(app.RegisterBean(cache.NewService()))
	neverror.PanicError(app.RegisterBean(cache.NewHandler()))

	app.Run()
}


