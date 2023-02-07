ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=SET
else
    SET=export
endif

all:
	go fmt $(foreach X,$(wildcard internal/*),&& pushd "$(X)" && go fmt && popd)
	go build

test:
	go test
	$(SET) "GOOS=linux" && $(MAKE) all
