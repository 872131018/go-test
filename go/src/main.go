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
		"martialarchery.com",
		"romancingthebrush.com",
	}

	// get which project to test on
	t, a := Prompt(p)

	// create a channel to hold agent results
	attempts := make(chan attempt)

	// generate a to make requests
	for i := 0; i < a; i++ {
		go newAgent(i, p[t], attempts)
	}

	// hash to send to display
	o := make(map[int]string)

	// listen to agents and aggregate results
	//go handleResponses(attempts, output)

	for a := range attempts {
		o[a.a] = fmt.Sprintf("Agent %d - successes: %d - failures: %d - duration: %v", a.a, a.s, a.e, a.l)
	}

	Display(o)
}

// func handleResponses(attempts chan attempt, output *map[int]string) {
// 	for a := range attempts {
// 		(*output)[a.agent] = fmt.Sprintf("Agent %d - successes: %d - failures: %d - duration: %v", a.agent, a.s, a.e, a.length)
// 	}
// }
