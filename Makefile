
.PHONY: test
test: # vanilla test
	go test ./test -race -count=1

.PHONY: test-timeout
test-timeout:
	go test ./test -bench=. -benchmem -race -timeout 30s -count=1

.PHONY: test-concurrent
test-concurrent:
	go test ./test -race -parallel 10 -count=1

.PHONY: build
build:
	go build -o app server/cmd/main.go

.PHONY: build-race
build-race:
	go build -race -o app server/cmd/main.go 

.PHONY: build-escape
build-escape:
	go build -gcflags "-m -l" -o app server/cmd/main.go

.PHONY: run
run:
	go build -o app cmd/main.go  && ./app

.PHONY: all_profile
all_profile:
	go test ./bench -run=none -bench=. -benchmem -benchtime=20s -memprofile=mem.pprof -cpuprofile=cpu.pprof -blockprofile=block.pprof

.PHONY: cpu_profile
cpu_profile:
	go test ./bench -bench=. -benchmem -cpuprofile=cpu.pprof

.PHONY: mem_profile
mem_profile:
	go test ./bench -bench=. -benchmem -memprofile=mem.pprof

.PHONY: cpu_profile-it
cpu_profile-it:
	go test ./bench -bench=. -benchmem -cpuprofile=cpu.pprof && go tool pprof cpu.pprof

.PHONY: mem_profile-it
mem_profile-it:
	go test ./bench -bench=. -benchmem -memprofile=mem.pprof && go tool pprof mem.pprof