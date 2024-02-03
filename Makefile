########################################################
# rebitcask

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
run: build
	./app

.PHONY: run-race
run-race: build-race
	./app

.PHONY: run-compose
run-compose: init-network
	docker-compose -f docker-compose.yml up -d

.PHONY: stop-compose
stop-compose:
	docker-compose -f docker-compose.yml stop

########################################################
# Profiling

.PHONY: all_profile
all_profile:
	go test ./bench -run=none -bench=. -benchmem -benchtime=20s -memprofile=mem.pprof -cpuprofile=cpu.pprof -blockprofile=block.pprof

.PHONY: cpu_profile
cpu_profile:
	go tool pprof -http=":8080" cpu.pprof

.PHONY: mem_profile
mem_profile:
	go tool pprof -http=":8080" mem.pprof

########################################################
# Chore

.PHONY: init
init: init-network
	go mod tidy && go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: init-network
init-network:
	@bash ./scripts/init-network.sh	

########################################################
# Discovery server commands

.PHONY: discovery-init
discovery-init: init-network
	go mod tidy && go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: discovery-build
discovery-build:
	cd discovery && swag init -g ./cmd/main.go -o ./docs && cd ../ && go build -o app.discovery ./discovery/cmd

.PHONY: discovery-run
discovery-run: discovery-build
	./app.discovery

.PHONY: discovery-run-compose
discovery-run-compose: init-network
	docker-compose -f docker-compose-discovery.yml up -d

.PHONY: discovery-stop-compose
discovery-stop-compose:
	docker-compose -f docker-compose-discovery.yml down