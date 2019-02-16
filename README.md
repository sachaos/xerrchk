# xerrchk

Static analysis tool for [xerrors](https://godoc.org/golang.org/x/xerrors).

This tool has not been fully tested. If you have problem in using this, please report issues or send PR in GitHub.

## Install

```sh
$ go get github.com/sachaos/xerrchk/cmd/xerrchk
```

## Features

* Find unwrapped error
* Find binary expression which is comparing error"
* Find invalid wrap format (using [tenntenn/gosa/passes/wraperrfmt](https://github.com/tenntenn/gosa/tree/master/passes/wraperrfmt))
* Checks returning nil when err is not nil (using [tenntenn/gosa/passes/nilerr](https://github.com/tenntenn/gosa/tree/master/passes/nilerr))
