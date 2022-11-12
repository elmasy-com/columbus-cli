# columbus-cli

A CLI application for Columbus.

See more at the [columbus-server](https://github.com/elmasy-com/columbus-server) repository.

**Lookup is free to use!**

Insert is requires an API key!. To use the API key, set the `COLUMBUS_KEY` environment variable.

> Currently, there is no way to get an API key. Thats a work in progress.
>
> I working on multiple solution to make it possible to contribute to the Columbus Database. 

## Install

```bash
wget -q 'https://github.com/elmasy-com/columbus-cli/releases/latest/download/columbus' -O columbus && \
wget -q 'https://github.com/elmasy-com/columbus-cli/releases/latest/download/columbus.sha' -O columbus.sha && \
sha512sum -c columbus.sha && rm columbus.sha && chmod +x columbus
```

## Usage

```bash
columbus help
```
```
Usage: ./columbus <command> <args>
To get more info about the command: ./columbus <command> help
Commands:
	lookup		Lookup domain
	insert		Insert domain
    help        Print this help

The API key must be set in COLUMBUS_KEY environment variable!
The server URI can be changed by setting the COLUMBUS_URI environment variable.
```

### Lookup

Do a lookup on the Columbus server and returns a newline sepratated list of every known hostname.

```bash
columbus lookup help
```
```
Usage: ./columbus lookup <domain>

Examples:
./columbus lookup example.com	-> Lookup for example.com

Returns a newline separated list of the full hostnames.
```

### Insert

Insert `<domain>` into the Columbus Database.

- On sucess, returns nothing with code 0. Duplications are silently ignored and count as a sucessful insert.
- On error, returns the error message and code 1.

```bash
columbus insert help
```
```
Usage: ./columbus insert <domain>

If <domain> is "input", then reads domains from the standard input.
If <domain> is "file <path>" then read domains from the given file file.

Examples:
echo 'example.com
www.example.com' | ./columbus insert input	-> Read and insert example.com and www.example.com
./columbus insert file /path/to/domains		-> Insert domains from the file
./columbus insert example.com			-> Insert example.com

IMPORTANT:
If "input" or "file" selected, than the domains must be newline separated (one domain per line).
This command requires a valid API key! Set API key in the COLUMBUS_KEY environment variable
```