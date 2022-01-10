package main

import (
	"bufio"
	"io"
	"log"
	"time"
)

func noThreads(file io.Reader) {
	log.Println("Starting with no threads..")
	start := time.Now()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// For each line (a single string), we send a message to an actor from a random pool
		s := scanner.Text()
		getHash(s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Println("Binomial took", elapsed)
}
