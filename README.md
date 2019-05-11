# Keylock

Golang keylock implement, used when user want to lock action by key,uuid or sth 

Easy and Fast

## Correctness

```
=== RUN   TestKeyMutex
--- PASS: TestKeyMutex (0.48s)
=== RUN   TestNewKeyLockWithLock
--- PASS: TestNewKeyLockWithLock (2.00s)
=== RUN   TestNewKeyLockWithoutLock
--- PASS: TestNewKeyLockWithoutLock (1.00s)
PASS
ok      klock   4.501s
```

## Performance
```
goos: darwin
goarch: amd64
pkg: klock
BenchmarkKeyLock-12      2000000               904 ns/op
```

## Compare To

```
goos: darwin
goarch: amd64
pkg: klock

# klock
BenchmarkKeyLock-12              2000000               797 ns/op              83 B/op          0 allocs/op

# github.com/alibaba/pouch/pkg/kmutex
BenchmarkPCKeyLock-12            1000000              1861 ns/op             322 B/op          2 allocs/op

# github.com/im7mortal/kmutex
BenchmarkKmutex-12                 10000            737253 ns/op             216 B/op          2 allocs/op

PASS
ok      klock   12.992s

```