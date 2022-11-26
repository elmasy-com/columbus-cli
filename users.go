package main

import (
	"fmt"
	"os"
	"strings"

	sdk "github.com/elmasy-com/columbus-sdk"
	"github.com/olekukonko/tablewriter"
)

func usersHelp() {

	fmt.Printf("USAGE\n")
	fmt.Printf("	%s users\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("RETURN\n")
	fmt.Printf("	Returns a table of every user.\n")
	fmt.Printf("	The key is redacted for security reason.\n")
	fmt.Printf("\n")
	fmt.Printf("EXAMPLE\n")
	fmt.Printf("	%s users\n", os.Args[0])
	fmt.Printf("	+---------------------+--------------+-------+\n")
	fmt.Printf("	|        NAME         |      KEY     | ADMIN |\n")
	fmt.Printf("	+---------------------+--------------+-------+\n")
	fmt.Printf("	| example1            | 1234****abcd | true  |\n")
	fmt.Printf("	| example2            | 1234****abcd | false |\n")
	fmt.Printf("	+---------------------+--------------+-------+\n")
}

func Users() {

	// More arguments than needed
	if len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Too much arguments: %s\n", strings.Join(os.Args[2:], " "))
		fmt.Fprintf(os.Stderr, "Use \"%s users help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	// Two argument, the third must be "help"
	if len(os.Args) == 3 {
		if os.Args[2] == "help" {
			usersHelp()
			os.Exit(0)
		}

		fmt.Fprintf(os.Stderr, "Only \"help\" can be used as an argument for users!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s users help\" to get help.\n", os.Args[0])
		os.Exit(1)
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

	users, err := sdk.GetUsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get users: %s\n", err)
		os.Exit(1)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Key", "Admin"})

	for i := range users {
		row := make([]string, 0)

		row = append(row, users[i].Name)
		row = append(row, fmt.Sprintf("%s****%s", users[i].Key[:4], users[i].Key[len(users[i].Key)-4:]))
		row = append(row, fmt.Sprintf("%v", users[i].Admin))

		table.Append(row)
	}

	table.Render()
}
