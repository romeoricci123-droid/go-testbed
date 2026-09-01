package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
)

const Version = "0.1.0"

func Greeting() string {
	return "hello from go-testbed"
}

func NewID() string {
	return uuid.New().String()
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
