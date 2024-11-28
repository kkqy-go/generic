package generic

import (
	"sync"
)

type SyncMap[KeyT, ValueT any] struct {
	syncMap sync.Map
}

func (p *SyncMap[KeyT, ValueT]) Load(key KeyT) (ValueT, bool) {
	value, ok := p.syncMap.Load(key)
	if ok {
		return value.(ValueT), ok
	} else {
		var T ValueT
		return T, ok
	}

}
func (p *SyncMap[KeyT, ValueT]) Store(key KeyT, value ValueT) {
	p.syncMap.Store(key, value)
}
func (p *SyncMap[KeyT, ValueT]) LoadOrStore(key KeyT, value ValueT) (ValueT, bool) {
	actual, loaded := p.syncMap.LoadOrStore(key, value)
	return actual.(ValueT), loaded
}

func (p *SyncMap[KeyT, ValueT]) LoadAndDelete(key KeyT) (ValueT, bool) {
	value, ok := p.syncMap.LoadAndDelete(key)
	if ok {
		return value.(ValueT), ok
	} else {
		var T ValueT
		return T, ok
	}
}

func (p *SyncMap[KeyT, ValueT]) Delete(key KeyT) {
	p.syncMap.Delete(key)
}
func (p *SyncMap[KeyT, ValueT]) Range(f func(key KeyT, value ValueT) bool) {
	p.syncMap.Range(func(key, value any) bool {
		return f(key.(KeyT), value.(ValueT))
	})
}

func (p *SyncMap[KeyT, ValueT]) Len() int {
	count := 0
	p.syncMap.Range(func(key, value any) bool {
		count += 1
		return true
	})
	return count
}
