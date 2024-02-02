
.PHONY: test
test: # vanilla test
	go test ./... -race -count=1

.PHONY: test-timeout
test-timeout:
	go test ./test -bench=. -benchmem -race -timeout 30s -count=1

.PHONY: test-concurrent
test-concurrent:
	go test ./test -race -parallel 10 -count=1

.PHONY: build
build:
	swag init -g ./cmd/main.go -o ./docs && go build -o app ./cmd

.PHONY: build-race
build-race:
	swag init -g ./cmd/main.go -o ./docs && go build -race -o app ./cmd

.PHONY: build-escape
build-escape:
	swag init -g ./cmd/main.go -o ./docs && go build -gcflags "-m -l" -o app ./cmd

.PHONY: run
run:
	swag init -g ./cmd/main.go -o ./docs && go build -o app ./cmd && ./app

.PHONY: run-race
run-race:
	swag init -g ./cmd/main.go -o ./docs && go build -race -o app ./cmd && ./app

.PHONY: run-prod
run:
	go build -o app ./cmd && ./app

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

.PHONY: init
init: init-network
	go mod tidy && go install github.com/swaggo/swag/cmd/swag@latest && swag init -g ./cmd/main.go -o ./docs

.PHONY: init-network
init-network:
	./init-network.sh	