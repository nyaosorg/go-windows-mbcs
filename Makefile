ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=SET
else
    SET=export
endif

all:
	go fmt $(foreach X,filter encoding $(wildcard internal/*),&& pushd "$(X)" && go fmt && popd)
	go build

test:
	go test $(foreach X,filter encoding,&& pushd $(X) && go test && popd)
	$(SET) "GOOS=linux" && $(MAKE) all
