GO = go
CFLAGS =
BIN = benri

all: build

build:
	$(GO) build $(CFLAGS) -o $(BIN)

install: build
	$(GO) install

test:
	$(GO) test -v ./...

