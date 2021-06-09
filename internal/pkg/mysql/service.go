// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xfali/gobatis"
	"github.com/xfali/neve-core/appcontext"
	"github.com/xfali/neve-example/internal/pkg/cache"
	"github.com/xfali/neve-example/internal/pkg/mysql/dao"
	"github.com/xfali/xlog"
)

type impl struct {
	log     xlog.Logger
	sessMgr *gobatis.SessionManager
}

func NewService() cache.Service {
	return &impl{
		log: xlog.GetLogger(),
	}
}

func (impl *impl) RegisterFunction(registry appcontext.InjectFunctionRegistry) error {
	return registry.RegisterInjectFunction(func(sessMgr *gobatis.SessionManager) {
		impl.sessMgr = sessMgr
	}, "testDb")
}

func (impl *impl) Get(key string) string {
	req := dao.TestTable{
		Key: key,
	}
	ret, err := req.Select(impl.sessMgr.NewSession())
	if err != nil {
		impl.log.Errorln(err)
	}
	if len(ret) == 0 {
		return ""
	}
	return ret[0].Value
}

func (impl *impl) Set(key, value string) {
	req := dao.TestTable{
		Key:   key,
		Value: value,
	}

	// Transaction
	impl.sessMgr.NewSession().Tx(func(session *gobatis.Session) error {
		i, err := req.Update(session)
		if err != nil {
			impl.log.Errorln(err)
			return err
		}
		// Not update, insert it.
		if i <= 0 {
			_, _, err := req.Insert(session)
			if err != nil {
				impl.log.Errorln(err)
			}
		}
		return err
	})
}

func (impl *impl) Delete(key string) bool {
	req := dao.TestTable{
		Key: key,
	}
	ret, err := req.Delete(impl.sessMgr.NewSession())
	if err != nil {
		impl.log.Errorln(err)
	}
	return ret > 0
}
