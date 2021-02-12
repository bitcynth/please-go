INSTALL_PATH := /usr/local/bin/please-go

.PHONY: build install

build: please-go

please-go:
	go build -o ./please-go

install: please-go
	cp please-go $(INSTALL_PATH)
	chown root:root $(INSTALL_PATH)
	chmod 4701 $(INSTALL_PATH)
