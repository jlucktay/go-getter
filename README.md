# go-getter

Fetch URLs over HTTP in parallel.

## Installation

### Prerequisites

You should have a [working Go environment](https://golang.org/doc/install) and have `$GOPATH/bin` in your `$PATH`.

### Compiling

To download the source, compile, and install the wrapper binary, run:

``` shell
go get github.com/jlucktay/go-getter/...
```

The source code will be located in `$GOPATH/src/github.com/jlucktay/go-getter/`.

A newly-compiled `go-getter` binary will be in `$GOPATH/bin/`.

## Usage

In a directory that also contains a `SampleListOfUrls.txt` file, run `go-getter`:

``` shell
$ go-getter
Content-Length results:
Minimum of 21,787 and maximum of 443,075 (total: 1,980,280) from 10 URLs.
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
