# Venlock (a.k.a. Vendor Lock)

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/crashdump/venlock/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/crashdump/venlock?status.svg)](https://godoc.org/github.com/crashdump/venlock)


## Install

```bash
go install github.com/crashdump/venlock/cmd/venlock@latest
```

## Use

```bash
venlock ./sources/
```

    ┌─────────────┐
    │ Vendor Lock │
    └─────────────┘
    
    NAME:
    venlock - Search for package manifests and identifies untrusted libraries.
    
    USAGE:
    venlock [global options] command [command options] [arguments...]
    
    AUTHOR:
    Adrien Pujol <ap@cdfr.net>
    
    COMMANDS:
    enumerate, e  enumerate all the libraries from source code.
    generate, g   generate a config.json from source code.
    enforce, v    enforce inventory libraries.
    help, h       Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
    --help  (default: false)


### Build

```bash
go build ./... -o dist/venlock
```

### Test

```bash
go test ./...
```
