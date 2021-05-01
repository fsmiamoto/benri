GO := $(shell which go)

BIN = benri
BENCH_COUNT=1000

all: build

build:
	$(GO) build -o $(BIN)

install: build
	sudo cp $(BIN) /usr/bin/$(BIN)

test:
	$(GO) test -v ./...

benchmark: build
	@echo "Running benri ${BENCH_COUNT} times..."
	@time -p ./benchmark.sh $(BIN) $(BENCH_COUNT) 2>&1 
