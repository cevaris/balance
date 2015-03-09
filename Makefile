all: build install

build:
	go build

install:
	go install

debug:
	go install -gcflags '-N'


