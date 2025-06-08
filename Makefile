ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=SET
    NUL=nul
    WHICH=where.exe
else
    SET=export
    NUL=/dev/null
    WHICH=which
endif

ifndef GO
    SUPPORTGO=go1.20.14
    GO:=$(shell $(WHICH) $(SUPPORTGO) 2>$(NUL) || echo go)
endif

all:
	$(GO) fmt ./...
	$(GO) build

get:
	$(GO) get -u
	$(GO) get golang.org/x/sys@v0.30.0
	$(GO) get golang.org/x/text@v0.22.0
	$(GO) mod tidy

test:
	$(GO) test
	$(SET) "GOOS=linux" && $(MAKE) all

bench:
	$(GO) test -bench . -benchmem
