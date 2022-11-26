package main

import (
	"fmt"
	"os"
	"strings"

	sdk "github.com/elmasy-com/columbus-sdk"
)

func userHelp() {

	fmt.Printf("USAGE\n")
	fmt.Printf("	%s user <target> <command>\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("TARGET\n")
	fmt.Printf("	self	<command>		Run command on self. Command can be \"name\", \"key\" or \"delete\".\n")
	fmt.Printf("	<name>	<command>		Run command on user with name <name>\n")
	fmt.Printf("	new	<name> <admin>		Create a new user with <name> and <admin>\n")
	fmt.Printf("	help				Print help\n")
	fmt.Printf("\n")
	fmt.Printf("COMMAND\n")
	fmt.Printf("	name <name>		Change the user's name to <name>\n")
	fmt.Printf("	key			Generate a new key for the user\n")
	fmt.Printf("	delete			Delete the user\n")
	fmt.Printf("	admin <value>		Set user's admin value to <value> (\"true\" or \"false\")\n")
	fmt.Printf("\n")
	fmt.Printf("NOTE\n")
	fmt.Printf("	Only admin user can use <name> target, regular user can only use \"self\".\n")
	fmt.Printf("	Only admin user can set admin value!\n")
	fmt.Printf("	If target is self, only name and key can be modified.\n")
	fmt.Printf("	Restricted usernames are \"self\", \"new\" and \"help\".\n")
	fmt.Printf("	\"delete\" will print the deleted user in case of mistake.\n")

}

func userSelf() {

	fmt.Fprintf(os.Stderr, "Not implemented!\n")
	os.Exit(2)
}

func userOther(u string) {

	if !sdk.DefaultUser.Admin {
		fmt.Fprint(os.Stderr, "Failed: your user is not admin!\n")
		os.Exit(1)
	}

	// Command
	switch os.Args[3] {
	case "name":
		if len(os.Args) < 5 {
			fmt.Fprintf(os.Stderr, "Name is missing!\n")
			os.Exit(1)
		}
		if len(os.Args) > 5 {
			fmt.Fprintf(os.Stderr, "Too much agument: %s\n", strings.Join(os.Args[5:], " "))
			os.Exit(1)
		}
		if os.Args[4] == "" {
			fmt.Fprintf(os.Stderr, "Name is empty\n")
			os.Exit(1)
		}

		target, err := sdk.GetOtherUser(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get %s: %s\n", os.Args[2], err)
			os.Exit(1)
		}

		err = sdk.ChangeOtherUserName(&target, os.Args[4])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to change name for %s: %s\n", target.Name, err)
			os.Exit(1)
		}
	case "key":

		target, err := sdk.GetOtherUser(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get %s: %s\n", os.Args[2], err)
			os.Exit(1)
		}

		err = sdk.ChangeOtherUserKey(&target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to change key for %s: %s\n", target.Name, err)
		}
		fmt.Printf("%s\n", target.Key)

	case "delete":

		target, err := sdk.GetOtherUser(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get %s: %s\n", os.Args[2], err)
			os.Exit(1)
		}

		err = sdk.Delete(target, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete %s: %s\n", target.Name, err)
			os.Exit(1)
		}

		// Print deleted user in case of mistake
		fmt.Printf("Name: %s\n", target.Name)
		fmt.Printf("Key: %s\n", target.Key)
		fmt.Printf("Amdin: %v\n", target.Admin)

	case "admin":
		if len(os.Args) < 5 {
			fmt.Fprintf(os.Stderr, "Value for admin is missing!\n")
			os.Exit(1)
		}

		value := false
		switch os.Args[4] {
		case "true":
			value = true
		case "false":
			value = false
		default:
			fmt.Fprintf(os.Stderr, "Invalid value for admin: %s\n", os.Args[4])
			os.Exit(1)
		}

		target, err := sdk.GetOtherUser(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get %s: %s\n", os.Args[2], err)
			os.Exit(1)
		}

		if target.Admin == value {
			fmt.Fprintf(os.Stderr, "Failed to set admin value for %s: value is already %v\n", target.Name, value)
			os.Exit(1)
		}

		err = sdk.ChangeOtherUserAdmin(&target, value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to set admin value for %s: %s\n", target.Name, err)
			os.Exit(1)
		}
	}
}

func userNew() {

	if !sdk.DefaultUser.Admin {
		fmt.Fprint(os.Stderr, "Failed to create user: your user is not admin!\n")
		os.Exit(1)
	}

	// Check restricted usernames
	if os.Args[3] == "self" || os.Args[3] == "new" || os.Args[3] == "help" {
		fmt.Fprintf(os.Stderr, "Restricted username: %s\n", os.Args[3])
		fmt.Fprintf(os.Stderr, "Restricted usernames are \"self\", \"new\" and \"help\"\n")
		os.Exit(1)
	}

	// Admin is missing (eg.: "columbus user new example")
	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Admin value for new user is missing!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s user help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	// Too much argument (eg.: "columbus user new test false invalid")
	if len(os.Args) > 5 {
		fmt.Fprintf(os.Stderr, "Too much argument: %s\n", strings.Join(os.Args[5:], " "))
		os.Exit(1)
	}

	admin := false

	switch os.Args[4] {
	case "true":
		admin = true
	case "false":
		admin = false
	default:
		fmt.Fprintf(os.Stderr, "Invalid user for admin: %s\n", os.Args[4])
		os.Exit(1)
	}

	u, err := sdk.AddUser(os.Args[3], admin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create user: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Key: %s\n", u.Key)
	fmt.Printf("Admin: %v\n", u.Admin)

}

func User() {

	// Fewer arguments than needed (eg.: "columbus user")
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Target for user is missing!\n")
		fmt.Fprintf(os.Stderr, "Use \"%s user help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	// Two argument (eg.: "columbus user example")
	if len(os.Args) < 4 {

		switch os.Args[2] {
		case "help":
			userHelp()
			os.Exit(0)
		case "new":
			fmt.Fprintf(os.Stderr, "Name for new user is missing!\n")
		default:
			fmt.Fprintf(os.Stderr, "Command for target is missing!\n")
		}

		fmt.Fprintf(os.Stderr, "Use \"%s user help\" to get help.\n", os.Args[0])
		os.Exit(1)
	}

	// Set server if needed
	if os.Getenv("COLUMBUS_URI") != "" {
		sdk.SetURI(os.Getenv("COLUMBUS_URI"))
	}

	if os.Getenv("COLUMBUS_KEY") == "" {
		fmt.Fprintf(os.Stderr, "COLUMBUS_KEY environment variable is empty!\n")
		os.Exit(1)
	}

	// Set DefaultUser
	err := sdk.GetDefaultUser(os.Getenv("COLUMBUS_KEY"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user: %s\n", err)
		os.Exit(1)
	}

	// Target
	switch os.Args[2] {
	case "self":
		userSelf()
	case "new":
		userNew()
	default:
		userOther(os.Args[3])
	}
}
