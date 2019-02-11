package main

import (
	"math"

	"github.com/gammazero/workerpool"
)

type SimpleTest struct {
	Url       string
	Steps     float64
	Samples   int
	Time      float64
	Rates     []float64
	ParaLimit int
}

func NewSimpleTest() *SimpleTest {
	return &SimpleTest{Rates: []float64{}}
}

func (t *SimpleTest) Run() {
	for _, rate := range t.Rates {
		wp := workerpool.New(math.MaxInt32)
		r := NewRatchet()
		s := int(t.Time * rate)
		for i := 0; i < s; i++ {
			p := newProbe(t.Url, rate)
			wp.Submit(func() {
				p.Run()
			})
			r.Wait(1 / rate)
		}
		wp.StopWait()
	}
}
