package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	sdk "github.com/elmasy-com/columbus-sdk"
	"github.com/elmasy-com/elnet/domain"
)

func insertHelp() {

	fmt.Printf("USAGE\n")
	fmt.Printf("	%s insert <domain>\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("	If <domain> is \"input\", then reads domains from the standard input.\n")
	fmt.Printf("	If <domain> is \"file <path>\" then read domains from the given file.\n")
	fmt.Printf("\n")
	fmt.Printf("INFO\n")
	fmt.Printf("If \"input\" or \"file\" selected, than the domains must be newline separated list (one domain per line).\n")
	fmt.Printf("This command requires a valid API key! Set API key in the COLUMBUS_KEY environment variable\n")
	fmt.Printf("\n")
	fmt.Printf("EXAMPLE\n")
	fmt.Printf("	echo 'example.com\\nwww.example.com' | %s insert input	-> Read and insert example.com and www.example.com\n", os.Args[0])
	fmt.Printf("	%s insert file /path/to/domains				-> Insert domains from the file\n", os.Args[0])
	fmt.Printf("	%s insert example.com					-> Insert example.com\n", os.Args[0])
}

// Insert domain(s) from input.
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

// Insert domain(s) from file in path.
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

// Insert a single domain.
func insertOne(d string) {

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

	// Fewer arguments than needed
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Domain for insert is missing!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s insert help\" to get help\n", os.Args[0])
		os.Exit(1)
	}

	// More arguments than needed
	if len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Too much arguments: %s\n", strings.Join(os.Args[2:], " "))
		fmt.Fprintf(os.Stderr, "Use \"%s lookup help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	// Able to print help before getting COLUMBUS_KEY
	if os.Args[2] == "help" {
		insertHelp()
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
		insertOne(os.Args[2])
	}
}
