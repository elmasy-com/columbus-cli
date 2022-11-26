package main

import (
	"fmt"
	"os"
	"strings"

	sdk "github.com/elmasy-com/columbus-sdk"
	"github.com/elmasy-com/elnet/domain"
)

func lookupHelp() {

	fmt.Printf("USAGE\n")
	fmt.Printf("	%s lookup <domain>\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("RETURN\n")
	fmt.Printf("	Returns a newline separated list of the full hostnames.\n")
	fmt.Printf("\n")
	fmt.Printf("EXAMPLE\n")
	fmt.Printf("	%s lookup example.com	-> Lookup for subdomains of example.com\n", os.Args[0])
}

func lookup(d string) {

	if !domain.IsValid(d) {
		fmt.Fprintf(os.Stderr, "Failed to lookup for %s: invalid domain\n", d)
		os.Exit(1)
	}

	if os.Getenv("COLUMBUS_URI") != "" {
		sdk.SetURI(os.Getenv("COLUMBUS_URI"))
	}

	subs, err := sdk.Lookup(d, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to lookup for %s: %s\n", d, err)
		os.Exit(1)
	}

	for i := range subs {
		if subs[i] == "" {
			fmt.Printf("%s\n", d)
			continue
		}
		fmt.Printf("%s.%s\n", subs[i], d)
	}
}

func Lookup() {

	// Fewer arguments than needed
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Domain for lookup is missing!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s lookup help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	// More arguments than needed
	if len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Too much arguments: %s\n", strings.Join(os.Args[2:], " "))
		fmt.Fprintf(os.Stderr, "Use \"%s lookup help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	if os.Args[2] == "help" {
		lookupHelp()
		os.Exit(0)
	}

	lookup(os.Args[2])
}
