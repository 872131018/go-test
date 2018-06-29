package main

import (
	"fmt"
	"net/http"
	"time"
)

func newAgent(id int, t string, attempts chan attempt) {
	// generate url to target
	t = fmt.Sprintf("http://%s/", t)

	// successes
	s := 0

	//errors
	e := 0

	// start testing on the target
	for i := 0; ; i++ {
		// request start
		now := time.Now()

		// make request to the target and track result
		res, err := http.Get(t)
		if err != nil {
			e++
		} else {
			s++
			res.Body.Close()
		}

		// data about the request and response
		a := attempt{
			a: id,
			s: s,
			e: e,
			l: time.Since(now),
		}

		// response ok get status code
		if res != nil {
			a.code = res.StatusCode
			a.err = nil
		}

		// response err track err
		if err != nil {
			a.code = 500
			a.err = err
		}

		// push attempt into channel
		attempts <- a

		// brief moment of silence
		time.Sleep(1000 * time.Millisecond)
	}
}
