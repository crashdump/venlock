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


```bash
venlock enumerate test/fixtures

```

    ┌─────────────┐
    │ Vendor Lock │
    └─────────────┘
    
    Enumerating libraries from source code...
    
    > Go...
    ... found 12 dependencies.
    
    github.com/PuerkitoBio/goquery
    github.com/avelino/slugify
    github.com/otiai10/copy
    github.com/yuin/goldmark
    golang.org/x/oauth2
    github.com/andybalholm/cascadia
    github.com/golang/protobuf
    golang.org/x/net
    golang.org/x/sys
    golang.org/x/text
    google.golang.org/appengine
    google.golang.org/protobuf
    
    > Maven...
    ... found 2 dependencies.
    
    foo:bar
    junit:junit
    
    > NPM...
    ... found 2 dependencies.
    
    express
    compression


### Build

```bash
go build ./... -o dist/venlock
```

### Test

```bash
go test ./...
```
