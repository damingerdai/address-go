init:
	chmod +x scripts/*.sh && sh scripts/init.sh

build: clean
	go build -o cmd/main cmd/address/main.go

run: build
	./cmd/main

migrate:
	rm -r -f ./build/migrate && sh scripts/migrate.sh

migrate-up: migrate
	./build/migrate up

migrate-down: migrate
	./build/migrate down

migrate-drop: migrate
	./build/migrate clear

docker:
	docker-compose down && docker-compose build address && docker-compose up

.PHONY: clean
clean:
	rm main || :