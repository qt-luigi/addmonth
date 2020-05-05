# adddate

addmonth adds months to a base date according to a format.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/addmonth
```

## Usage

```sh
$ addmonth -h
	-m  2 -b 20200229 -> 20200430
	-m -1 -b 20200331 -> 20200229

Usage:
	adddate [-m months]
	        [-b basedate | -f format | -b basedate -f format]
Arguments are:
	-m months
		adding months
	-b basedate
		base date (default: today (ex. 20200505))
		When don't use with -f option, a format is yyyymmdd
	-f format
		base date format by Go (default: 20060102)
		When don't use with -b option, a value is 20060102
```

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)
