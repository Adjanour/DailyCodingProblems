package main

import (
 "github.com/adjanour/DailyCodingProblems/day10"
	"flag"
	"fmt"
	"os"

	"github.com/adjanour/DailyCodingProblems/day09"

	"github.com/adjanour/DailyCodingProblems/day08"
)

// auto-inserted imports go here

var registry = map[string]func(){
   "10": day10.Run,
	"09": day09.Run,
	"08": day08.Run,
	// auto-inserted entries go here
}

func main() {
	day := flag.String("day", "", "Specify the day to run (e.g. 01, 02)")
	flag.Parse()

	if run, ok := registry[*day]; ok {
		run()
	} else {
		fmt.Println("Invalid or missing day. Use --day=XX")
		os.Exit(1)
	}
}
