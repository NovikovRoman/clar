BUILD_DIR=bin

build:
	mkdir -p ${BUILD_DIR}
	go build -ldflags="-s -w" -o ${BUILD_DIR}/clar