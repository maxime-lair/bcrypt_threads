package main

// go get golang.org/x/crypto/bcrypt

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func getHash(s string) {

	toHashString := []byte(s)
	_, err := bcrypt.GenerateFromPassword(toHashString, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	file, err := os.Open("data/string_list_full.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	noThreads(file)
	withThreads(file)

}
