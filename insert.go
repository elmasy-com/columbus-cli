package main

import (
	"bufio"
	"fmt"
	"os"

	sdk "github.com/elmasy-com/columbus-sdk"
	"github.com/elmasy-com/elnet/domain"
)

func InsertHelp() {

	fmt.Printf("Usage: %s insert <domain>\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("If <domain> is \"input\", then reads domains from the standard input.\n")
	fmt.Printf("If <domain> is \"file <path>\" then read domains from the given file file.\n")
	fmt.Printf("\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("echo 'example.com\nwww.example.com' | %s insert input	-> Read and insert example.com and www.example.com\n", os.Args[0])
	fmt.Printf("%s insert file /path/to/domains		-> Insert domains from the file\n", os.Args[0])
	fmt.Printf("%s insert example.com			-> Insert example.com\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("IMPORTANT:\n")
	fmt.Printf("If \"input\" or \"file\" selected, than the domains must be newline separated (one domain per line).\n")
	fmt.Printf("This command requires a valid API key! Set API key in the COLUMBUS_KEY environment variable\n")
}

func insertInput() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		if !domain.IsValid(scanner.Text()) {
			fmt.Fprintf(os.Stderr, "Failed to insert %s: invalid domain\n", scanner.Text())
			continue
		}

		err := sdk.Insert(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to insert %s: %s\n", scanner.Text(), err)
			os.Exit(1)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "Failed to read from stdin: %s\n", scanner.Err())
		os.Exit(1)
	}
}

func insertFile(path string) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open %s: %s\n", path, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if !domain.IsValid(scanner.Text()) {
			fmt.Fprintf(os.Stderr, "Failed to insert %s: invalid domain\n", scanner.Text())
			continue
		}

		err := sdk.Insert(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to insert %s: %s\n", scanner.Text(), err)
			os.Exit(1)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "Failed to read from %s: %s\n", path, scanner.Err())
		os.Exit(1)
	}
}

func insert(d string) {

	if !domain.IsValid(d) {
		fmt.Fprintf(os.Stderr, "Failed to insert %s: invalid domain\n", d)
		os.Exit(1)
	}

	err := sdk.Insert(d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert: %s\n", err)
		os.Exit(1)
	}
}

func Insert() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Domain for insert is missing!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s insert help\" to get help\n", os.Args[0])
		os.Exit(1)
	}

	if os.Args[2] == "help" {
		InsertHelp()
		os.Exit(0)
	}
	if os.Getenv("COLUMBUS_URI") != "" {
		sdk.SetURI(os.Getenv("COLUMBUS_URI"))
	}

	if os.Getenv("COLUMBUS_KEY") == "" {
		fmt.Fprintf(os.Stderr, "COLUMBUS_KEY environment variable is empty!\n")
		os.Exit(1)
	}

	err := sdk.GetDefaultUser(os.Getenv("COLUMBUS_KEY"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user: %s\n", err)
		os.Exit(1)
	}

	switch os.Args[2] {
	case "input":
		insertInput()
	case "file":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "File path is missing!\n")
			os.Exit(1)
		}
		insertFile(os.Args[3])
	default:
		insert(os.Args[2])
	}
}
