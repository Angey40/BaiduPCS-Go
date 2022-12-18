package cachemap

import (
	"github.com/iikira/BaiduPCS-Go/baidupcs/expires"
	"sync"
)

var (
	GlobalCacheOpMap = CacheOpMap{}
)

type (
	CacheOpMap struct {
		cachePool sync.Map
	}
)

func (cm *CacheOpMap) LazyInitCachePoolOp(op string) CacheUnit {
	cacheItf, ok := cm.cachePool.Load(op)
	if !ok {
		cache := &cacheUnit{}
		cm.cachePool.Store(op, cache)
		return cache
	}
	return cacheItf.(CacheUnit)
}

// ClearInvalidate 清除已过期的数据(一般用不到)
func (cm *CacheOpMap) ClearInvalidate() {
	cm.cachePool.Range(func(_, cacheItf interface{}) bool {
		cache := cacheItf.(CacheUnit)
		cache.Range(func(key interface{}, exp expires.DataExpires) bool {
			if exp.IsExpires() {
				cache.Delete(key)
			}
			return true
		})
		return true
	})
}

// PrintAll 输出所有缓冲项目
func (cm *CacheOpMap) PrintAll() {

}
