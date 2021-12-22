BINARY_NAME=main
API_PATH=cmd/api/main.go

.PHONY: build
build:
	go build -o ${BINARY_NAME} ${API_PATH}

.PHONY: run
run:
	./${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm -f ${BINARY_NAME}