package main

import (
	"fmt"
	"os"

	sdk "github.com/elmasy-com/columbus-sdk"
	"github.com/elmasy-com/elnet/domain"
)

func lookupHelp() {

	fmt.Printf("Usage: %s lookup <domain>\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("%s lookup example.com	-> Lookup for example.com\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("Returns a newline separated list of the full hostnames.\n")
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

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Domain for lookup is missing!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s lookup help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}
	switch os.Args[2] {
	case "help":
		lookupHelp()
	default:
		lookup(os.Args[2])
	}
}
