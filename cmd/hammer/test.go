package main

import (
	"time"
)

type Test interface {
	Run()
}

var t0 = time.Now()
