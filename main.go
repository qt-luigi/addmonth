package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	defdatefmt = "20060102"
	usage      = `addmonth adds months to a base date according to a format.
When a base date is last day of the month, a result is also last day of the month.

ex.
	-m  2 -b 20200229 -> 20200430
	-m -1 -b 20200331 -> 20200229

Usage:
	adddate [-m months]
	        [-b basedate | -f format | -b basedate -f format]
Arguments are:
	-m months
		adding months
	-b basedate
		base date (default: today (ex. %s))
		When don't use with -f option, a format is yyyymmdd
	-f format
		base date format by Go (default: %s)
		When don't use with -b option, a value is %s
`
)

var (
	months   int
	basedate string
	format   string
)

func init() {
	now := time.Now().Format(defdatefmt)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, now, defdatefmt, defdatefmt)
	}
	flag.IntVar(&months, "m", 0, "adding months")
	msgBasedate := fmt.Sprintf("base date (default: today (ex. %s))", now)
	flag.StringVar(&basedate, "b", now, msgBasedate)
	msgFormat := fmt.Sprintf("base date format by Go (default: %s)", defdatefmt)
	flag.StringVar(&format, "f", defdatefmt, msgFormat)
}

func main() {
	flag.Parse()
	d, err := time.Parse(format, basedate)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	s := AddMonth(d, months).Format(format)
	if len(s) != len(format) {
		fmt.Fprintf(os.Stderr, "%s exceeds a format digit\n", s)
		os.Exit(1)
	}
	fmt.Println(s)
}

// AddMonth returns the time corresponding to adding the
// given number of months to t.
// For example, AddMonth(t, 2) applied to January 1, 2011
// (= t) returns March 1, 2011.
//
// AddMonth does not normalize its result in the same way
// that Date does, so, for example, adding one month to
// October 31 yields November 30.
func AddMonth(t time.Time, months int) time.Time {
	lastMonthDay := func(t time.Time) int {
		return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Day()
	}

	// Creating 1st Date from t and adding months because AddDate() normalizes t.
	am := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, months, 0)
	ad := lastMonthDay(am)

	if ld := lastMonthDay(t); t.Day() == ld || t.Day() > ad {
		return time.Date(am.Year(), am.Month(), ad, am.Hour(), am.Minute(), am.Second(), am.Nanosecond(), am.Location())
	}

	return t.AddDate(0, months, 0)
}
