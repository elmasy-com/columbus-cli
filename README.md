# columbus-cli

A CLI application for Columbus.

See more at the [columbus-server](https://github.com/elmasy-com/columbus-server) repository.

**Lookup is free to use!**

Insert is requires an API key!. To use the API key, set the `COLUMBUS_KEY` environment variable.

> Currently, there is no way to get an API key. Thats a work in progress.
>
> I working on multiple solution to make it possible to contribute to the Columbus Project. 

## Install

The checksum file is signed with key `10BC80B36072944B5678AF395D00FD9E9F2A3725`.
```bash
gpg --receive-key 10BC80B36072944B5678AF395D00FD9E9F2A3725
```

On Linux/amd64:

```bash
wget -q 'https://github.com/elmasy-com/columbus-cli/releases/latest/download/columbus-linux-amd64' -O columbus-linux-amd64 && \
wget -q 'https://github.com/elmasy-com/columbus-cli/releases/latest/download/checksums' -O checksums && \
gpg --verify checksums && sha512sum --ignore-missing -c checksums && rm checksums && \
mv columbus-linux-amd64 columbus && chmod +x columbus
```

### Build

Requirements:
- `go 1.19+`

Clone the repository and build with the standard build command:
```bash
go build -o columbus .
```

## Usage

```bash
columbus help
```
```
USAGE
	./columbus <command> <subcommand>

COMMANDS
	lookup		Lookup domain
	insert		Insert domain
	user		Run command on user
	users		Get a list of every user
	version		Print version
	help		Print this help

INFO
	To get more info about the command: ./columbus <command> help.
	The API key must be set in COLUMBUS_KEY environment variable!
	The server URI can be changed by setting the COLUMBUS_URI environment variable.
```

### Lookup

Do a lookup on the Columbus server and returns a newline sepratated list of every known hostname.

```bash
columbus lookup help
```
```
USAGE
	./columbus lookup <domain>

RETURN
	Returns a newline separated list of the full hostnames.

EXAMPLE
	./columbus lookup example.com	-> Lookup for subdomains of example.com
```

### Insert

Insert `<domain>` into the Columbus Database.

- On sucess, returns nothing with code 0. Duplications are silently ignored and count as a sucessful insert.
- On error, returns the error message and code 1.

When using `input` or `file`, an invalid domain does not stop the process, only print the error to stderr.

```bash
columbus insert help
```
```
USAGE
	./columbus insert <domain>

	If <domain> is "input", then reads domains from the standard input.
	If <domain> is "file <path>" then read domains from the given file.

INFO
If "input" or "file" selected, than the domains must be newline separated list (one domain per line).
This command requires a valid API key! Set API key in the COLUMBUS_KEY environment variable

EXAMPLE
	echo 'example.com\nwww.example.com' | ./columbus insert input	-> Read and insert example.com and www.example.com
	./columbus insert file /path/to/domains				-> Insert domains from the file
	./columbus insert example.com					-> Insert example.com
```

### User operations

```bash
columbus user help
```
```
USAGE
	./columbus user <target> <command>

TARGET
	self	<command>		Run command on self. Command can be "name", "key" or "delete".
	<name>	<command>		Run command on user with name <name>
	new	<name> <admin>		Create a new user with <name> and <admin>
	help				Print help

COMMAND
	name <name>		Change the user's name to <name>
	key			Generate a new key for the user
	delete			Delete the user
	admin <value>		Set user's admin value to <value> ("true" or "false")

NOTE
	Only admin user can use <name> target, regular user can only use "self".
	Only admin user can set admin value!
	If target is self, only name and key can be modified.
	Restricted usernames are "self", "new" and "help".
	"delete" will print the deleted user in case of mistake.
```

#### Users

Return every user on the server.

```bash
columbus users help
```
```
USAGE
	./columbus users

RETURN
	Returns a table of every user.
	The key is redacted for security reason.

EXAMPLE
	./columbus users
	+---------------------+--------------+-------+
	|        NAME         |      KEY     | ADMIN |
	+---------------------+--------------+-------+
	| example1            | 1234****abcd | true  |
	| example2            | 1234****abcd | false |
	+---------------------+--------------+-------+
```