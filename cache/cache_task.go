package cache

import (
	"golly/app"
	"sync"
)

type TaskCache struct {
	lock  sync.Mutex
	cache map[int]*app.Link
}

func (tc *TaskCache) Cache() map[int]*app.Link {
	return tc.cache
}

func NewTaskCache() *TaskCache {
	return &TaskCache{
		cache: make(map[int]*app.Link),
	}
}

func (tc *TaskCache) Put(l *app.Link) {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	tc.cache[l.ID] = l
}

func (tc *TaskCache) Get(id int) *app.Link {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	return tc.cache[id]
}

func (tc *TaskCache) Delete(id int) {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	delete(tc.cache, id)
}
