build:
	mkdir -p bin
	go build -ldflags="-s -w" -o bin/clar