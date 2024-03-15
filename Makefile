########################################################
# rebitcask

.PHONY: gen_grpc
gen_grpc:
	protoc --go-grpc_out=. --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative server/chorepb/*.proto server/rebitcaskpb/*.proto

.PHONY: test
test: # vanilla test
	go test ./... -race -count=1 -v

.PHONY: test-timeout
test-timeout:
	go test ./test -bench=. -benchmem -race -timeout 30s -count=1 -v

.PHONY: test-concurrent
test-concurrent:
	go test ./test -race -parallel 10 -count=1 -v

.PHONY: build
build:
	swag init -g ./cmd/main.go -o ./docs --exclude ./discovery && go build -o app ./cmd

.PHONY: build-race
build-race:
	swag init -g ./cmd/main.go -o ./docs --exclude ./discovery && go build -race -o app ./cmd

.PHONY: build-escape
build-escape:
	swag init -g ./cmd/main.go -o ./docs --exclude ./discovery && go build -gcflags "-m -l" -o app ./cmd

.PHONY: run
run: build
	./app

.PHONY: run-race
run-race: build-race
	./app

.PHONY: compose-up
compose-up: init-network
	docker-compose -f docker-compose.yml up -d

.PHONY: compose-down
compose-down:
	docker-compose -f docker-compose.yml down

.PHONY: compose-build
compose-build:
	docker-compose -f docker-compose.yml build

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
init:
	go mod tidy && go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: init-network
init-network:
	@bash ./scripts/init-network.sh	

########################################################
# Discovery server commands

# .PHONY: init-discovery
# init-discovery:
# 	go mod tidy && go install github.com/swaggo/swag/cmd/swag@latest

# .PHONY: build-discovery
# build-discovery:
# 	cd discovery && swag init -g ./cmd/main.go -o ./docs && cd ../ && go build -o app.discovery ./discovery/cmd

# .PHONY: run-discovery
# run-discovery: build-discovery
# 	./app.discovery

.PHONY: init-discovery
init-discovery:
	go mod tidy

.PHONY: build-discovery
build-discovery:
	go build -o app.discovery ./discovery/cmd

.PHONY: run-discovery
run-discovery: build-discovery
	./app.discovery

.PHONY: compose-up-discovery
compose-up-discovery: init-network
	docker-compose -f docker-compose-discovery.yml up -d

.PHONY: compose-down-discovery
compose-down-discovery:
	docker-compose -f docker-compose-discovery.yml down

.PHONY: compose-build-discovery
compose-build-discovery: 
	docker-compose -f docker-compose-discovery.yml build 

.PHONY: compose-build-up-discovery
compose-build-up-discovery: compose-build-discovery
	make compose-up-discovery


################################################################
# Run cluster
.PHONY: cluster-up
cluster-up:
	make compose-up-discovery && make compose-up

.PHONY: cluster-down
cluster-down:
	make compose-down-discovery && make compose-down

.PHONY: cluster-build
cluster-build:
	make compose-build-discovery && make compose-build


################################################################
# Generate Grpc code
.PHONY: grpc-gen
grpc-gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative server/*/*.proto

