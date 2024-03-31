.Phony: default clean
BINARY_NAME=./normalizer/normalizerapp


default:
	go build -o=${BINARY_NAME} ./normalizer

clean:
	go clean
	rm ${BINARY_NAME}
