build: clean
	go build -o cmd/main main.go

run: build
	./main 
.PHONY: clean
clean:
	rm main || :