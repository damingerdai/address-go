clean:
	rm main || :

build: clean
	go build main.go

run: build
	./main 