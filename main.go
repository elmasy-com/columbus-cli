package main

import (
	"fmt"
	"os"
)

func Help() {

	fmt.Printf("USAGE\n")
	fmt.Printf("	%s <command> <subcommand>\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("COMMANDS\n")
	fmt.Printf("	lookup		Lookup domain\n")
	fmt.Printf("	insert		Insert domain\n")
	fmt.Printf("	user		Run command on user\n")
	fmt.Printf("	users		Get a list of every user\n")
	fmt.Printf("	version		Print version\n")
	fmt.Printf("	help		Print this help\n")
	fmt.Printf("\n")
	fmt.Printf("INFO\n")
	fmt.Printf("	To get more info about the command: %s <command> help.\n", os.Args[0])
	fmt.Printf("	The API key must be set in COLUMBUS_KEY environment variable!\n")
	fmt.Printf("	The server URI can be changed by setting the COLUMBUS_URI environment variable.\n")
}

var (
	Version string
	Commit  string
)

func main() {

	if len(os.Args) < 2 {
		Help()
		os.Exit(1)
	}

	// Command
	switch os.Args[1] {
	case "lookup":
		Lookup()
	case "insert":
		Insert()
	case "user":
		User()
	case "users":
		Users()
	case "version":
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Commit: %s\n", Commit)
	case "help":
		Help()
	default:
		fmt.Fprintf(os.Stderr, "Unkown command: %s\n", os.Args[1])
		fmt.Printf("Use \"%s help\" to get help\n", os.Args[0])
		os.Exit(1)
	}
}
