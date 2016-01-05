package main

import "fmt"

var (
	version   = "pre-release"
	buildDate = "unknown"
	gitCommit = "unknown"
)

func printVersion() {
	fmt.Println("Version: ", version)
	fmt.Println("Build date: ", buildDate)
	fmt.Println("Git commit: ", gitCommit)
}
