# Lib Guardian

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/crashdump/libguardian/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/crashdump/libguardian?status.svg)](https://godoc.org/github.com/crashdump/libguardian)


## Install

```bash
go install github.com/crashdump/libguardian/cmd/libguardian@latest
```

## Use

```bash
libguardian ./sources/
```

    ┌──────────────┐
    │ Lib Guardian │
    └──────────────┘

    > Searching for supported manifests...
      Found foo/bar/package.json.

    > Examining files...
      Found foreign library: foo.bar
      Found foreign library: bar.baz

    Failed!


### Build

```bash
go build ./... -o dist/libguardian
```

### Test

```bash
go test ./...
```
