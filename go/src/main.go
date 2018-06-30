package main

import (
	"fmt"
	"time"
)

type attempt struct {
	a    int
	s    int
	e    int
	l    time.Duration
	code int
	err  error
}

func main() {
	// define possible ts
	p := []string{
		"localhost:8080",
		"martialarchery.com",
		"romancingthebrush.com",
	}

	// get which project to test on
	t, a := Prompt(p)

	// create a channel to hold agent results
	attempts := make(chan attempt)

	// request counter
	rps := 0

	// create output map to display results
	o := make(map[int]string)

	// generate a to make requests
	for i := 0; i < a; i++ {
		go newAgent(i, p[t], attempts)
	}

	// write results to central place
	go func(attempts chan attempt) {
		for a := range attempts {
			o[a.a] = fmt.Sprintf("Agent %d - successes: %d - failures: %d - duration: %v", a.a, a.s, a.e, a.l)
		}
	}(attempts)

	// listen to a and aggregate results
	go Handle(attempts, &o, &rps)

	Display(o, &rps)
}
