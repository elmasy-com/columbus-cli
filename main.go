package main

import (
	"fmt"
	"os"
)

func HelpPrint() {

	fmt.Printf("Usage: %s <command> <args>\n", os.Args[0])
	fmt.Printf("To get more info about the command: %s <command> help\n", os.Args[0])
	fmt.Printf("Commands:\n")
	fmt.Printf("	lookup		Lookup domain\n")
	fmt.Printf("	insert		Insert domain\n")
	fmt.Printf("	help		Print this help\n")
	fmt.Printf("\n")
	fmt.Printf("The API key must be set in COLUMBUS_KEY environment variable!\n")
	fmt.Printf("The server URI can be changed by setting the COLUMBUS_URI environment variable.\n")
}

var (
	Version string
	Commit  string
)

func main() {

	if len(os.Args) == 1 {
		HelpPrint()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "lookup":
		Lookup()
	case "insert":
		Insert()
	case "help":
		HelpPrint()
	case "version":
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Commit: %s\n", Commit)
	default:
		fmt.Fprintf(os.Stderr, "Unkown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
