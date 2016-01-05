package main

import "fmt"

var (
	version   = "Prerelease"
	buildDate = "Unknown"
	gitCommit = "Unknown"
)

func printVersion() {
	fmt.Println("Version: ", version)
	fmt.Println("Build date: ", buildDate)
	fmt.Println("Git commit: ", gitCommit)
}
