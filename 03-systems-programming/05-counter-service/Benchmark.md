## Benchmark output

```
➜  concurrency-2-b-prework go test -bench=.
goos: darwin
goarch: arm64
pkg: counterservice
BenchmarkCounterServices/unsynchronized-8               442669327                2.703 ns/op
BenchmarkCounterServices/atomic-8                       131857444                9.076 ns/op
BenchmarkCounterServices/mutex-8                        64507600                18.60 ns/op
BenchmarkCounterServices/channel-8                       3895717               307.8 ns/op
BenchmarkCounterServicesContended/unsynchronized-8              1000000000               0.9907 ns/op
BenchmarkCounterServicesContended/atomic-8                      16075904                74.20 ns/op
BenchmarkCounterServicesContended/mutex-8                        7616943               151.9 ns/op
BenchmarkCounterServicesContended/channel-8                      3657214               313.1 ns/op
PASS
ok      counterservice  13.359s

```
