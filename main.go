// Main package
// 6 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package main

import (
	"github.com/KaiserGald/tph/util/cli"
)

// CLI Flags
var (
	w         bool
	newWorker bool
	userName  string
	pass      string
)

func processCLI() {
	processFlags()
	processArgs()
	handleCLI()
}

func processFlags() {
	cli.CreateCommand(&userName, "u", "Gald", "this is for the username")
}

func processArgs() {

}

func handleCLI() {

}

func main() {
	processCLI()

}
