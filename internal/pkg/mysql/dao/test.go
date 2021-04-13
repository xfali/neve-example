// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package dao

import (
	"github.com/xfali/gobatis"
)

func SelectTestTable(sess *gobatis.Session, model TestTable) ([]TestTable, error) {
	var dataList []TestTable
	err := sess.Select("test.selectTestTable").Param(model).Result(&dataList)
	return dataList, err
}

func SelectTestTableCount(sess *gobatis.Session, model TestTable) (int64, error) {
	var ret int64
	err := sess.Select("test.selectTestTableCount").Param(model).Result(&ret)
	return ret, err
}

func InsertTestTable(sess *gobatis.Session, model TestTable) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("test.insertTestTable").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func InsertBatchTestTable(sess *gobatis.Session, models []TestTable) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("test.insertBatchTestTable").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func UpdateTestTable(sess *gobatis.Session, model TestTable) (int64, error) {
	var ret int64
	err := sess.Update("test.updateTestTable").Param(model).Result(&ret)
	return ret, err
}

func DeleteTestTable(sess *gobatis.Session, model TestTable) (int64, error) {
	var ret int64
	err := sess.Delete("test.deleteTestTable").Param(model).Result(&ret)
	return ret, err
}
