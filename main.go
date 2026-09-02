package main

import (
	"flag"
	"fmt"
)

const Version = "0.1.0"

func Greeting() string {
	return "hello from go-testbed"
}

func Farewell() string {
	return "goodbye from go-testbed"
}

func main() {
	version := flag.Bool("version", false, "print version")
	quiet := flag.Bool("quiet", false, "suppress output")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		return
	}
	if !*quiet {
		fmt.Println(Greeting())
	}
}
