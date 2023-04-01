## main
`./main -cpuprofile cpu.prof`
`go tool pprof -cum -top cpu.prof`
`go tool pprof -png cpu.prof`
`go tool pprof -http=:8888 cpu.prof`
`curl -s http://localhost:8888/debug/pprof/profile > cpu.prof`


## calc, hsd
`go test`
`go test ./...`
`go test -short`
`go test -v -cover`
## something
`go test -count 5 -benchmem -bench . 2>&1 | tee old.log`
`go install golang.org/x/perf/cmd/benchstat@latest`
`benchstat old.log new.log`

Copyright (c) mattn
Released under the MIT license

https://github.com/mattn/go-hsd