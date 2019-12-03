build: clean
	go build -o cmd/main main.go

run: build
	./cmd/main 
.PHONY: clean
clean:
	rm main || :