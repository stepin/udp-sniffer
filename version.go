package main

import "fmt"

var (
	// nolint
	version = "pre-release"
	// nolint
	date = "unknown"
	// nolint
	commit = "unknown"
)

func printVersion() {
	fmt.Println("Version: ", version)
	fmt.Println("Build date: ", date)
	fmt.Println("Git commit: ", commit)
}
