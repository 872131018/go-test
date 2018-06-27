package main

import (
	"fmt"
	"time"
)

type attempt struct {
	agent  int
	s      int
	e      int
	length time.Duration
	code   int
	err    error
}

func main() {
	// define possible targets
	projects := []string{
		"martialarchery.com",
		"romancingthebrush.com",
	}

	// get which project to test on
	target, agents := Prompt(projects)

	// create a channel to hold agent results
	attempts := make(chan attempt)

	// generate agents to make requests
	for i := 0; i < agents; i++ {
		go newAgent(i, projects[target], attempts)
	}

	// hash to send to display
	output := make(map[int]string)

	// listen to agents and aggregate results
	go handleResponses(attempts, &output)

	Display(output)
}

func handleResponses(attempts chan attempt, output *map[int]string) {
	for a := range attempts {
		(*output)[a.agent] = fmt.Sprintf("Agent %d - successes: %d - failures: %d - duration: %v", a.agent, a.s, a.e, a.length)
	}
}
