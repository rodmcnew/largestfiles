package largestfiles

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Item struct {
	IsDir      bool
	Path       string
	Size       int64
	ChildCount int64
}

type ScanOptions struct {
	Path                   string
	Count                  int
	IgnoreFileSystemErrors bool
}

type ScanResult struct {
	Dirs           []Item
	Files          []Item
	TotalDirCount  int
	TotalFileCount int
	TotalSize      float64
}

type OnScanProgressUpdate func(size float64, fileCount int, dirCount int)

// Scan the file system for the largest files and directories.
//
// Note: This code assumes that filepath.Walk() can call its callback in many concurrent goroutines even though this may not be true.
//       This is to ensure future compatability with concurrent file system walkers, such as third party walkers.
func ScanFileSys(options ScanOptions, onProgressUpdate OnScanProgressUpdate, progressUpdateInterval time.Duration) (*ScanResult, error) {
	// This channel handles file system items as they are read by the walker and then stored in memory
	itemsChan := make(chan Item, 1)

	// Handle file system items by dumping the item's data into the itemsChan channel.
	readFileInfo := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if options.IgnoreFileSystemErrors {
				log.Println(err)
				return nil
			} else {
				return err
			}
		}
		itemsChan <- Item{IsDir: info.IsDir(), Path: path, Size: info.Size()}
		return nil
	}

	// Walk the file system
	var walkError error = nil
	go func() {
		err := filepath.Walk(options.Path, readFileInfo)
		if err != nil {
			if !options.IgnoreFileSystemErrors {
				walkError = err
			}
		}
		close(itemsChan)
	}()

	files := make([]Item, 0)              // Stores a list of files
	dirs := make([]Item, 0)               // Stores a list of directories
	pathToDirIMap := make(map[string]int) // Used to quickly look up directories during size calculation
	var totalSize float64 = 0             // Stores the total disk usage found

	// Start a ticker that will display our progress reading the filesystem every so often
	progressDisplayTicker := time.NewTicker(progressUpdateInterval)
	go func() {
		for range progressDisplayTicker.C {
			onProgressUpdate(totalSize, len(files), len(dirs))
		}
	}()

	// As the file path walker reads the file system, store the data it returns
	for item := range itemsChan {
		if item.IsDir {
			dirs = append(dirs, item)
			pathToDirIMap[item.Path] = len(dirs) - 1
		} else {
			files = append(files, item)
			totalSize += float64(item.Size)
		}
	}
	if walkError != nil {
		return nil, walkError
	}

	// We are done reading the filesystem so stop the display-progress ticker
	progressDisplayTicker.Stop()

	// For each file, add its size to its parent directory
	for _, item := range files {
		dir := &dirs[pathToDirIMap[filepath.Dir(item.Path)]]
		dir.Size += item.Size
		dir.ChildCount++
	}

	// Sort the files and directories by size
	sort.Slice(files, func(i, j int) bool { return files[i].Size > files[j].Size })
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Size > dirs[j].Size })

	// Slice out any files or directories that are beyond the display length
	totalFileCount := len(files)
	if totalFileCount > options.Count {
		files = files[:options.Count]
	}
	totalDirCount := len(dirs)
	if totalDirCount > options.Count {
		dirs = dirs[:options.Count]
	}

	return &ScanResult{
		Dirs:           dirs,
		Files:          files,
		TotalFileCount: totalFileCount,
		TotalDirCount:  totalDirCount,
		TotalSize:      totalSize,
	}, nil
}
