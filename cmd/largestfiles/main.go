package main

import (
	"os"
	"text/tabwriter"
	"time"

	"github.com/rodmcnew/largestfiles/pkg/largestfiles"
	"github.com/rodmcnew/largestfiles/third_party/bytecount"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Controls how often we display a status update to the user while scanning the file system
const printProgressUpdateInterval = time.Second / 2

func main() {
	// Create a locale-aware printer so that numbers get formatted with thousands seperators
	p := message.NewPrinter(language.English)

	// Prints the size and counts on the same line we are already on
	printStatus := func(size float64, fileCount int, dirCount int) {
		p.Printf("\rFound %s of usage in %d files and %d directories.", bytecount.ByteCountDecimal(int64(size)), fileCount, dirCount)
	}

	// Parse the CLI args into options
	options := cliArgsToOptions()

	// Print the path that we are scanning inside
	p.Printf("Looking in %s\n", options.Path)

	// Scan the file system while printing the status every so often
	scanStartTime := time.Now()
	result, err := largestfiles.ScanFileSys(options, printStatus, printProgressUpdateInterval)
	if err != nil {
		p.Println("\n", err)
		os.Exit(1)
	}

	// Print the status one last time with the final values
	printStatus(result.TotalSize, result.TotalFileCount, result.TotalDirCount)
	p.Printf(" Scanning took %s\n", time.Since(scanStartTime).Round(time.Millisecond))

	// Print a table of the largest directories
	p.Println("\n- - - - - - - - - - Largest Directories - - - - - - - - - -")
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 0, ' ', tabwriter.AlignRight)
	p.Fprintln(writer, p.Sprintf("Size  \tFiles  \tPath"))
	for _, item := range result.Dirs {
		p.Fprintln(writer, p.Sprintf("%s  \t%d  \t%s", bytecount.ByteCountDecimal(int64(item.Size)), item.ChildCount, item.Path))
	}
	writer.Flush()

	// Print a table of the largest files
	p.Println("\n- - - - - - - - - - - Largest Files - - - - - - - - - - - -")
	writer = tabwriter.NewWriter(os.Stdout, 0, 4, 0, ' ', tabwriter.AlignRight)
	p.Fprintln(writer, p.Sprintf("Size  \tPath"))
	for _, file := range result.Files {
		p.Fprintln(writer, p.Sprintf("%s  \t%s", bytecount.ByteCountDecimal(int64(file.Size)), file.Path))
	}
	writer.Flush()
}
