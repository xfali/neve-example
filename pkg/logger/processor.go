// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package logger

import (
	"github.com/xfali/fig"
	"github.com/xfali/neve-core/bean"
	"github.com/xfali/xlog"
	"github.com/xfali/xlog/writer"
	"io"
	"os"
)

type processor struct {
}

func NewLoggerProcessor() *processor {
	return &processor{}
}

func (p *processor) Init(conf fig.Properties, container bean.Container) error {
	path := conf.Get("logger.path", "./neve.log")
	xlog.SetOutput(io.MultiWriter(os.Stdout, writer.NewBufferedRotateFileWriter(&writer.BufferedRotateFile{
		Path:            path,
		RotateFrequency: writer.RotateEveryDay,
		RotateFunc:      writer.ZipLogsAsync,
	})))
	return nil
}

func (p *processor) Classify(o interface{}) (bool, error) {
	return false, nil
}

func (p *processor) Process() error {
	return nil
}

func (p *processor) BeanDestroy() error {
	return nil
}
