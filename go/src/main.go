package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type success struct {
	id    int
	s     int
	e     int
	start time.Time
	code  int
}

type failure struct {
	err   error
	id    int
	e     int
	start time.Time
}

func main() {
	var wg sync.WaitGroup
	successes := make(chan success)
	failures := make(chan failure)

	agents, err := strconv.Atoi(os.Getenv("AGENTS"))
	if err != nil {
		log.Fatal("Must have a valid number of agents")
	}
	for i := 0; i < agents; i++ {
		wg.Add(1)
		go makeRequestAgent(successes, failures, &wg)
	}
	go printResponses(successes)
	go printErrors(failures)
	wg.Wait()
}

func printResponses(successes chan success) {
	for res := range successes {
		fmt.Printf("Agent %d successes: %d\n", res.id, res.s)
		fmt.Printf("Last request duration: %v\n", time.Since(res.start))
	}
}
func printErrors(failures chan failure) {
	for err := range failures {
		fmt.Printf("Agent %d errors: %d\n", err.id, err.e)
		fmt.Printf("Last request duration: %v\n", time.Since(err.start))
	}
}

func makeRequestAgent(successes chan success, errors chan failure, wg *sync.WaitGroup) {
	defer wg.Done()
	id := rand.Int()
	s := 0
	e := 0
	for i := 0; ; i++ {
		now := time.Now()
		res, err := http.Get("http://martialarchery.com/")
		if err != nil {
			e++
			errors <- failure{
				err:   err,
				id:    id,
				e:     e,
				start: now,
			}
			next := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(next)
		} else {
			s++
			successes <- success{
				id:    id,
				s:     s,
				e:     e,
				start: now,
				code:  res.StatusCode,
			}
			res.Body.Close()
			next := time.Duration(rand.Intn(2)) * time.Second
			time.Sleep(next)
		}
	}
}
