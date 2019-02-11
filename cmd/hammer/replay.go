package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
)

type ReplayElement struct {
	Wait time.Duration
	Url string
}

type ReplayTest struct {
	Elements []ReplayElement
}

func NewReplayTest(file string) *ReplayTest {
	var rt ReplayTest
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &rt.Elements)
	if err != nil {
		panic(err)
	}
	return &rt
}

func (t *ReplayTest) Run() {
	t0 = time.Now()
	for _, e := range t.Elements {
		t1 := t0.Add(e.Wait)
		delay := time.Until(t1)
		p := newProbe(e.Url, delay.Seconds())
		go p.Run()
		fmt.Fprintf(os.Stderr, "%fs until next request\n", delay.Seconds())
		time.Sleep(delay)
	}
}
