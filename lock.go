package klock

import (
	"sync"
)

type keyLock struct {
	localLockMap map[string]*lock
	globalLock   sync.Mutex
}

type lock struct {
	mux      *sync.Mutex
	refCount int
}

func NewKeyLock() *keyLock {
	return &keyLock{localLockMap: map[string]*lock{}}
}

func (km *keyLock) Lock(key string) {
	km.globalLock.Lock()

	wl, locked := km.localLockMap[key]

	if !locked {
		wl = &lock{
			mux:      new(sync.Mutex),
			refCount: 0,
		}
		km.localLockMap[key] = wl
	}

	wl.refCount++

	km.globalLock.Unlock()

	wl.mux.Lock()
}

func (km *keyLock) Unlock(key string) {
	km.globalLock.Lock()

	wl, locked := km.localLockMap[key]

	if !locked {
		km.globalLock.Unlock()
		return
	}

	wl.refCount--

	if wl.refCount <= 0 {
		delete(km.localLockMap, key)
	}

	km.globalLock.Unlock()

	wl.mux.Unlock()
}
