package main

import (
	"fmt"
	"net/http"
	"time"
)

func newAgent(id int, target string, attempts chan attempt) {
	// generate url to target
	target = fmt.Sprintf("http://%s/", target)

	// successes
	s := 0

	//errors
	e := 0

	// start testing on the target
	for i := 0; ; i++ {
		// request start
		now := time.Now()

		// make request to the target and track result
		res, err := http.Get(target)
		if err != nil {
			e++
		} else {
			s++
			res.Body.Close()
		}

		// data about the request and response
		attempt := attempt{
			agent:  id,
			s:      s,
			e:      e,
			length: time.Since(now),
		}

		// response ok get status code
		if res != nil {
			attempt.code = res.StatusCode
			attempt.err = nil
		}

		// response err track err
		if err != nil {
			attempt.code = 500
			attempt.err = err
		}

		// push attempt into channel
		attempts <- attempt

		// brief moment of silence
		time.Sleep(1000 * time.Millisecond)
	}
}
