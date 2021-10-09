package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rodmcnew/largestfiles/pkg/largestfiles"
)

// Controls how many files and directories we display when the user doesn't set this with a CLI arg
const defaultDisplayCount = 20

func cliArgsToOptions() largestfiles.ScanOptions {
	// The path is either the current directory or the first arg that doesn't start with "-"
	path := "./"
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			path = arg
			break
		}
	}

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: largestfiles [directory to scan]\n")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(flag.CommandLine.Output(), " -%s %s\n", f.Name, f.Usage)
		})
	}

	// Parse CLI flag args
	countPtr := flag.Int("c", defaultDisplayCount, "Number of items to display")
	ignoreFileSystemErrorsPtr := flag.Bool("i", false, "Ignore file system errors")
	flag.Parse()

	return largestfiles.ScanOptions{
		Path:                   path,
		Count:                  *countPtr,
		IgnoreFileSystemErrors: *ignoreFileSystemErrorsPtr,
	}
}
