package main

import (
	"math"
	"log"
	"strings"
	"time"
	"strconv"
)

func lerp(a, b, t float64) float64 {
	t = math.Min(math.Max(t, 0), 1)
	if b < a {
		t = 1 - t
		b, a = a, b
	}
	return t * (b - a) + a
}

type Ratchet struct {
	t0 time.Time
	t1 time.Time
}

func NewRatchet() *Ratchet {
	t := time.Now()
	return &Ratchet{t0: t, t1: t}
}

func (r *Ratchet) Wait(t float64) {
	r.t1 = r.t1.Add(time.Duration(t * float64(time.Second)))
	time.Sleep(time.Until(r.t1))
}

func parseRate (str string) float64 {
	parts := strings.Split(str, "/")
	if len(parts) == 1 {
		return 1 / parseDuration(parts[0])
	} else if len(parts) == 2 {
		num := parseFloat64(parts[0])
		dur := parseDuration(parts[1])
		return num / dur
	} else {
		log.Panicf("bad rate: %s", str)
		return 0
	}
}

func parseInt (str string) int {
	num, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		log.Panicf("bad number: %s", str)
	}
	return int(num)
}

func parseFloat64 (str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Panicf("bad number: %s", str)
	}
	return num
}

func parseDuration (str string) float64 {
	dur, err := time.ParseDuration(str)
	if err != nil {
		dur, err = time.ParseDuration("1" + str)
	}
	if err != nil {
		log.Panicf("bad duration: %s", str)
	}
	return dur.Seconds()
}
