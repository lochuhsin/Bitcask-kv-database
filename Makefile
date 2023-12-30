
.PHONY: test
test: # vanilla test
	go test ./test -bench=. -benchmem -race

.PHONY: test-timeout
test-timeout:
	go test ./test -bench=. -benchmem -race -timeout 30s

.PHONY: test-concurrent
test-concurrent:
	go test ./test -bench=. -benchmem -race -parallel 10

.PHONY: build
build:
	go build cmd/main.go

.PHONY: build-escape
build-escape:
	go build -gcflags "-m -l" -o app cmd/main.go

.PHONY: run
run:
	go build -o app cmd/main.go  && ./app

.PHONY: all_profile
all_profile:
	go test ./test -bench=. -run=none -benchmem -memprofile=mem.pprof -cpuprofile=cpu.pprof

.PHONY: cpu_profile
cpu_profile:
	go test ./test -bench=. -benchmem -cpuprofile=cpu.pprof

.PHONY: mem_profile
mem_profile:
	go test ./test -bench=. -benchmem -memprofile=mem.pprof

.PHONY: cpu_profile-it
cpu_profile-it:
	go test ./test -bench=. -benchmem -cpuprofile=cpu.pprof && go tool pprof cpu.pprof

.PHONY: mem_profile-it
mem_profile-it:
	go test ./test -bench=. -benchmem -memprofile=mem.pprof && go tool pprof mem.pprof