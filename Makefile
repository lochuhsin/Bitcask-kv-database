
.PHONY: test
test:
 	go test ./test -bench=. -benchmem

.PHONY: build
build:
	go build cmd/main.go

.PHONY: build_escape
build_escape:
	go build -gcflags "-m -l" -o app cmd/main.go

.PHONY: run
run:
	go build -o app cmd/main.go  && ./app

.PHONY: gen_all_profile
all_profile:
	go test ./test -bench=. -run=none -benchmem -memprofile=mem.pprof -cpuprofile=cpu.pprof

.PHONY: cpu_profile
cpu_profile:
	go test ./test -bench=. -benchmem -cpuprofile=cpu.pprof && go tool pprof cpu.pprof

.PHONY: mem_profile
mem_profile:
	go test ./test -bench=. -benchmem -memprofile=mem.pprof && go tool pprof mem.pprof
