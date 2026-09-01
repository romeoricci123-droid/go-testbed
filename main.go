package main

import (
	"flag"
	"fmt"
	"strings"
)

const Version = "0.1.0"

func Greeting() string {
	return "hello from go-testbed"
}

func FormatGreeting(uppercase bool) string {
	greeting := Greeting()
	if uppercase {
		return strings.ToUpper(greeting)
	}
	return greeting
}

func main() {
	version := flag.Bool("version", false, "print version")
	quiet := flag.Bool("quiet", false, "suppress output")
	uppercase := flag.Bool("uppercase", false, "print greeting in uppercase")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		return
	}
	if !*quiet {
		fmt.Println(FormatGreeting(*uppercase))
	}
}
