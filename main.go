package main

import "os"

// These variables are set in build step
var (
	Version = "unset"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
