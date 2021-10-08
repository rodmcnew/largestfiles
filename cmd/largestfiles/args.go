package main

import (
	"flag"
	"os"
	"strings"

	"github.com/rodmcnew/largestfiles/pkg/largestfiles"
)

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
