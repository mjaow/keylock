package klock

import (
	"sync"
	"testing"

	pcmutex "github.com/alibaba/pouch/pkg/kmutex"
	km "github.com/im7mortal/kmutex"
)

func TestKeyMutex(t *testing.T) {
	keyMutex := NewKeyLock()

	var count = 0

	var wg sync.WaitGroup

	var num = 10000

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keyMutex.Lock("a")
			count += i
			keyMutex.Unlock("a")
		}(i)
	}

	wg.Wait()

	expected := 50005000

	if count != expected {
		t.Fatalf("exptected %d and actual %d", expected, count)
	}
}

func BenchmarkKeyLock(b *testing.B) {
	keyMutex := NewKeyLock()

	var count = 0

	var wg sync.WaitGroup

	for i := 0; i <= b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keyMutex.Lock("a")
			count += i
			keyMutex.Unlock("a")
		}(i)
	}

	wg.Wait()
}

func BenchmarkPCKeyLock(b *testing.B) {
	keyMutex := pcmutex.New()

	var count = 0

	var wg sync.WaitGroup

	for i := 0; i <= b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keyMutex.Lock("a")
			count += i
			keyMutex.Unlock("a")
		}(i)
	}

	wg.Wait()
}

func BenchmarkKmutex(b *testing.B) {
	keyMutex := km.New()

	var count = 0

	var wg sync.WaitGroup

	for i := 0; i <= b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keyMutex.Lock("a")
			count += i
			keyMutex.Unlock("a")
		}(i)
	}

	wg.Wait()
}
