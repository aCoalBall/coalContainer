CMD=go
BIN_PATH=bin 
SRC_PATH=src 

all: clean build install

build:
	$(CMD) build -o $(BIN_PATH)/coalContainer $(SRC_PATH)/*

install:
	cp bin/coalContainer /usr/bin/coalContainer
	cp bin/coalContainer /usr/local/bin/coalContainer
	mkdir -p /var/lib/coalContainer/images
	mkdir -p /var/lib/coalContainer/volumes
	mkdir -p /var/lib/coalContainer/containers

uninstall:
	rm -rf /usr/bin/coalContainer /usr/local/bin/coalContainer
	rm -rf bin/coalContainer

clean: uninstall