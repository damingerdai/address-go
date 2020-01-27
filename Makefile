init:
	chmod +x scripts/*.sh && sh scripts/init.sh

build: clean
	go build -o cmd/main cmd/address/main.go

run: build
	./cmd/main

.PHONY: clean
clean:
	rm main || :