hello:
	echo "hello"

build:
	go mod init github.com/maxime-lair/bcrypt_threads
	go mod tidy
	go get golang.org/x/crypto/bcrypt
	
run:
	go run .

clean:
	go clean

all: build run

