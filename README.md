# golang lock by key

Easy and Fast

## Correctness

```
=== RUN   TestKeyMutex
--- PASS: TestKeyMutex (0.46s)
PASS
ok      klock   1.476s
```

## Performance
```
goos: darwin
goarch: amd64
pkg: klock
BenchmarkKeyLock-12      2000000               904 ns/op
PASS
ok      klock   2.964s
```