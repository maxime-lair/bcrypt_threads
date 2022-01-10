BINARY_NAME=bcrypt_threads

hello:
	echo "hello"

build:
	go get golang.org/x/crypto/bcrypt
	go build -o $(BINARY_NAME) main.go

run:
	./$(BINARY_NAME)

clean:
	go clean
	rm $(BINARY_NAME)

all: build run

