GO = go
CFLAGS =
BIN = benri
BENCH_COUNT=1000

all: build

build:
	$(GO) build $(CFLAGS) -o $(BIN)

install: build
	cp $(BIN) /usr/bin/$(BIN)

test:
	$(GO) test -v ./...

benchmark: build
	@echo "Running benri ${BENCH_COUNT} times..."
	@time -p ./benchmark.sh $(BIN) $(BENCH_COUNT) 2>&1 
