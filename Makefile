all: install

install:
	go install

debug:
	go install -gcflags '-N'


