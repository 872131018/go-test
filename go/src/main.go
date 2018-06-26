package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
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
	// create a channel to hold agent results
	attempts := make(chan attempt)
	// determine the number of agents to produce
	agents, err := strconv.Atoi(os.Getenv("AGENTS"))
	if err != nil {
		log.Fatal("Must have a valid number of agents")
	}
	// generate agents to make requests
	for i := 0; i < agents; i++ {
		go makeRequestAgent(i, attempts)
	}
	// hash to send to display
	output := make(map[int]string)
	// listen to agents and aggregate results
	go handleResponses(attempts, &output)
	// refresh display on an interval
	ticker := time.NewTicker(100 * time.Millisecond)
	for _ = range ticker.C {
		writeOutput(output)
	}
}

func handleResponses(attempts chan attempt, output *map[int]string) {
	for a := range attempts {
		(*output)[a.agent] = fmt.Sprintf("Agent %d - successes: %d - failures: %d - duration: %v", a.agent, a.s, a.e, a.length)
	}
}

func writeOutput(output map[int]string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("") //hack to fix output

	var keys []int
	for k := range output {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for key := range keys {
		fmt.Println(output[key])
	}
}

func makeRequestAgent(id int, attempts chan attempt) {
	s := 0
	e := 0
	for i := 0; ; i++ {
		now := time.Now()
		res, err := http.Get(os.Getenv("TARGET"))
		if err != nil {
			e++
		} else {
			s++
			res.Body.Close()
		}
		attempt := attempt{
			agent:  id,
			s:      s,
			e:      e,
			length: time.Since(now),
		}
		if res != nil {
			attempt.code = res.StatusCode
			attempt.err = nil
		}
		if err != nil {
			attempt.code = 500
			attempt.err = err
		}
		attempts <- attempt
	}
}
