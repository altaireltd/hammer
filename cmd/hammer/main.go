package main

import (
	"fmt"
	"strings"

	"github.com/alexflint/go-arg"
)

type Args struct {
	URL       string `arg:"-u,required"`
	Time      string `arg:"-t"`
	Rates     string `arg:"-r"`
	ParaLimit int    `arg:"-j"`
	Loop      bool   `arg:"-l"`
}

var args Args

func parseRates(rates string) []float64 {
	result := []float64{}
	for _, rate := range strings.Split(rates, ",") {
		result = append(result, parseRate(rate))
	}
	return result
}

func main() {
	args = Args{
		Time:      "10s",
		Rates:     "100ms",
		ParaLimit: 100,
	}
	arg.MustParse(&args)
	st := NewSimpleTest()
	st.ParaLimit = args.ParaLimit
	st.Url = args.URL
	st.Time = parseDuration(args.Time)
	st.Rates = parseRates(args.Rates)
	fmt.Println("start,stop,duration,status,size,hash,simul,rate")
start:
	st.Run()
	if args.Loop {
		goto start
	}
}
