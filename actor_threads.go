package main

import (
	"bufio"
	"io"
	"log"
	"sync"
	"time"
)

func actor(id int, toHashChannel chan string, timeToStop chan bool) {

	// Limit our actor to work every 200 ms
	for range time.Tick(100 * time.Millisecond) {
		select {
		case s := <-toHashChannel:
			getHash(s)

		default:
			// received nothing, check for stop
			//fmt.Println("[", id, "] No more job, stopping")
			select {
			case timeToStop <- true:
				//log.Println("[", id, "] Advertising to coworkers to stop")
			default:
				//log.Println("[", id, "] No stop message sent")
			}
			return
		}
	}
}

func withThreads(file io.Reader) {
	log.Println("Starting with threads..")
	start := time.Now()

	scanner := bufio.NewScanner(file)

	toHashChannel := make(chan string)
	timeToStop := make(chan bool)
	var wg sync.WaitGroup

	// Create 5 actors max
	max_actor := 5
	wg.Add(5)
	for i := 0; i < max_actor; i++ {
		go func(j int) {
			defer wg.Done()
			actor(j, toHashChannel, timeToStop)
		}(i)
	}

	for scanner.Scan() {
		// For each line (a single string), we send a message to an actor from a random pool
		s := scanner.Text()
		go func(toHash string) {
			toHashChannel <- s // send a string on the channel and wait for an actor to take it
		}(s)
	}

	// Wait for our pool of actors to finish nicely
	wg.Wait()
	//log.Println("Closing channels.")
	close(toHashChannel)
	close(timeToStop)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Println("Binomial took", elapsed)
}
