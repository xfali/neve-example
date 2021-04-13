// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package cache

import "sync"

type Service interface {
	Get(key string) string
	Set(key, value string)
	Delete(key string) bool
}

type impl struct {
	cache sync.Map
}

func NewService() Service {
	return &impl{
	}
}

func (impl *impl) Get(key string) string {
	v, ok := impl.cache.Load(key)
	if ok {
		return v.(string)
	}
	return ""
}

func (impl *impl) Set(key, value string) {
	impl.cache.Store(key, value)
}

func (impl *impl) Delete(key string) bool {
	_, ok := impl.cache.Load(key)
	impl.cache.Delete(key)
	return ok
}
