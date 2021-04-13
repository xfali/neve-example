// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package dao

import (
	"github.com/xfali/gobatis"
	"time"
)

type TestTable struct {
	//TableName gobatis.ModelName `test_table`
	Id         int       `xfield:"id"`
	Key        string    `xfield:"key"`
	Value      string    `xfield:"value"`
	Createtime time.Time `xfield:"createtime"`
}

func (m *TestTable) Select(sess *gobatis.Session) ([]TestTable, error) {
	return SelectTestTable(sess, *m)
}

func (m *TestTable) Count(sess *gobatis.Session) (int64, error) {
	return SelectTestTableCount(sess, *m)
}

func (m *TestTable) Insert(sess *gobatis.Session) (int64, int64, error) {
	return InsertTestTable(sess, *m)
}

func (m *TestTable) Update(sess *gobatis.Session) (int64, error) {
	return UpdateTestTable(sess, *m)
}

func (m *TestTable) Delete(sess *gobatis.Session) (int64, error) {
	return DeleteTestTable(sess, *m)
}
