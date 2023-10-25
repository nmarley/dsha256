.PHONY: clean build strip default

default: build strip

build:
	tinygo build

strip: dsha256
	llvm-strip dsha256

clean:
	go clean
