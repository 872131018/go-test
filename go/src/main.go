package main

import (
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

	// generate a to make requests
	for i := 0; i < a; i++ {
		go newAgent(i, p[t], attempts)
	}

	// hash to send to display
	o := make(map[int]string)

	// listen to a and aggregate results
	go Handle(attempts, &o, &rps)

	Display(o, &rps)
}
