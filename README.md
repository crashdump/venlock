# Venlock (a.k.a. Vendor Lock)

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/crashdump/venlock/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/crashdump/venlock?status.svg)](https://godoc.org/github.com/crashdump/venlock)


## Install/Run

You can download precompiled binaries, containers or install directly from source.

### Binaries (arm, amd64)

Precompiled binaries can be [found here](https://github.com/crashdump/venlock/releases).

### Docker (arm, amd64)

```bash
docker run -ti docker pull ghcr.io/crashdump/venlock:latest 
```

### Source

```bash
go install github.com/crashdump/venlock/cmd/venlock@latest
```

## Use

### Subcommands

```bash
./venlock
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


### Enumerate

```bash
./venlock enumerate test/fixtures
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


### Enforce

```bash
./venlock enforce -c test/fixtures/config.json test/fixtures
```

    enforce -c test/fixtures/config.json test/fixtures
    ┌─────────────┐
    │ Vendor Lock │
    └─────────────┘
    
    Searching for foreign libraries in source code...
    
    > Go...
    ... found foreign libraries:
      - github.com/PuerkitoBio/goquery
      - github.com/yuin/goldmark
      - golang.org/x/oauth2
      - github.com/andybalholm/cascadia
      - github.com/golang/protobuf
      - golang.org/x/net
      - golang.org/x/sys
      - golang.org/x/text
      - google.golang.org/appengine
      - google.golang.org/protobuf
    
    > Maven...
    
    No mismatch.
    
    > Npm...
    
    No mismatch.
    
    non-compliant: found unexpected libraries


## Contribute

### Build

```bash
go build ./... -o dist/venlock
```

### Test

```bash
go test ./...
```
