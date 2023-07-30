ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=SET
else
    SET=export
endif

all:
	go fmt
	go build

test:
	go test
	$(SET) "GOOS=linux" && $(MAKE) all

bench:
	go test -bench . -benchmem
