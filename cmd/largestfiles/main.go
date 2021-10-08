package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/rodmcnew/largestfiles/pkg/largestfiles"
	"github.com/rodmcnew/largestfiles/third_party/bytecount"
)

const printProgressUpdateInterval = time.Second / 2

func main() {
	// Parse the CLI args into options
	options := cliArgsToOptions()

	// Print the path that we are scanning inside
	fmt.Printf("Looking in %s\n", options.Path)

	// Scan the file system while printing the status every so often
	scanStartTime := time.Now()
	result, err := largestfiles.ScanFileSys(options, printStatus, printProgressUpdateInterval)
	if err != nil {
		fmt.Println("\n", err)
		os.Exit(1)
	}

	// Print the status one last time with the final values
	printStatus(result.TotalSize, result.TotalFileCount, result.TotalDirCount)
	fmt.Printf(" Scanning took %s\n", time.Since(scanStartTime).Round(time.Millisecond))

	// Print a table of the largest directories
	fmt.Println("\n- - - - - - - - - - Largest Directories - - - - - - - - - -")
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 0, ' ', tabwriter.AlignRight)
	fmt.Fprintln(writer, fmt.Sprintf("Size  \tFiles  \tPath"))
	for _, item := range result.Dirs {
		fmt.Fprintln(writer, fmt.Sprintf("%s  \t%d  \t%s", bytecount.ByteCountDecimal(int64(item.Size)), item.ChildCount, item.Path))
	}
	writer.Flush()

	// Print a table of the largest files
	fmt.Println("\n- - - - - - - - - - - Largest Files - - - - - - - - - - - -")
	writer = tabwriter.NewWriter(os.Stdout, 0, 4, 0, ' ', tabwriter.AlignRight)
	fmt.Fprintln(writer, fmt.Sprintf("Size  \tPath"))
	for _, file := range result.Files {
		fmt.Fprintln(writer, fmt.Sprintf("%s  \t%s", bytecount.ByteCountDecimal(int64(file.Size)), file.Path))
	}
	writer.Flush()
}

// Prints the size and counts on the same line we are already on
func printStatus(size float64, fileCount int, dirCount int) {
	fmt.Printf("\rFound %s of usage in %d files and %d directories.", bytecount.ByteCountDecimal(int64(size)), fileCount, dirCount)
}
