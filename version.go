package main

import "fmt"

var (
	version = "pre-release"
	date    = "unknown"
	commit  = "unknown"
)

func printVersion() {
	fmt.Println("Version: ", version)
	fmt.Println("Build date: ", date)
	fmt.Println("Git commit: ", commit)
}
