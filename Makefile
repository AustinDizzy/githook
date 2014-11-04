all: build

build: install_deps
	@go build
	@mkdir -p bin
	@mv githook bin/

clean:
	@rm -rf bin/*
	@rm -rf log/*
	@rm -rf pid/*

install: build install_service
	sudo cp bin/githook /usr/local/bin/githook
	sudo cp githook.json /usr/local/etc/githook.json

install_deps:
	@go get

install_service:
	sudo cp init.d/githook /etc/init.d/githook
	sudo chmod +X /etc/init.d/githook
