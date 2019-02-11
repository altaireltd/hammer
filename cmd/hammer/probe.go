package main

import (
	"crypto/tls"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"
)

type Probe struct {
	rate float64
	c    *http.Client
	r    *http.Request
}

var live int32

var tlsConfig = &tls.Config{InsecureSkipVerify: true}
var transport = &http.Transport{TLSClientConfig: tlsConfig}
var client = &http.Client{
	Timeout:   10 * time.Second,
	Transport: transport,
}

func newProbe(url string, rate float64) *Probe {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	return &Probe{rate, client, req}
}

type result struct {
	d0     time.Duration
	d1     time.Duration
	status string
	size   int
	hash   uint32
	live   int
	rate   float64
}

func (p *Probe) run() (r result) {
	r.rate = p.rate
	r.d0 = time.Since(t0)
	r.d1 = r.d0
	liveNow := atomic.AddInt32(&live, 1)
	defer atomic.AddInt32(&live, -1)
	r.live = int(liveNow - 1)
	if liveNow > int32(args.ParaLimit) {
		r.status = "SKIPPED"
		return
	}
	res, err := p.c.Do(p.r)
	r.d1 = time.Since(t0)
	if err == nil {
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err == nil {
			r.size = len(data)
			h := fnv.New32()
			h.Write(data)
			r.hash = h.Sum32()
			r.status = fmt.Sprint(res.StatusCode)
		} else {
			r.status = "FAILED_READ"
		}
	} else {
		r.status = "UNKNOWN_ERROR"
		if err.(*url.Error).Timeout() {
			r.status = "REQUEST_TIMEOUT"
		}
	}
	return
}

func (p *Probe) Run() {
	r := p.run()
	fmt.Printf("%f,%f,%f,%s,%d,%d,%d,%f\n",
		r.d0.Seconds(),
		r.d1.Seconds(),
		(r.d1 - r.d0).Seconds(),
		r.status,
		r.size,
		r.hash,
		r.live,
		r.rate)
}
