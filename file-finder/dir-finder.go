package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"sort"
)

// go run dir-finder.go ./test_dir

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Please prover valid dir\n")
		return
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Printf("Error during reading dir: %s. %v", args[0], err)
		return
	}

	// ---- Optimization: allocation memory upfront for empty files ----
	//var emptyFiles []byte          // Unoptimized
	//totalMem := len(files)*256     // Too much memory allocated
	totalMem := optimizedMemoryAllocation(files)
	emptyFiles := make([]string, 0, totalMem)

	for _, el := range files {
		info, _ := el.Info()

		if info.Size() == 0 {
			i := len(emptyFiles)
			lineToAdd := fmt.Sprintf("%d. %s", i+1, info.Name())
			emptyFiles = append(emptyFiles, lineToAdd)
			continue
		}

		if info.IsDir() {
			fmt.Printf("Inner Dir name: \"%s\". Size: %d Bytes \n", info.Name(), info.Size())
		}

		fmt.Printf("File name: \"%s\". Size: %d Bytes \n", info.Name(), info.Size())

	}

	if len(emptyFiles) != 0 {
		sort.Strings(emptyFiles)
		fmt.Printf("\nEmpty files: %s", emptyFiles)

		buf := &bytes.Buffer{}
		enErr := gob.NewEncoder(buf).Encode(emptyFiles)
		if enErr != nil {
			fmt.Printf("Error during encoding: %v", enErr)
			return
		}
		bs := buf.Bytes()

		writeErr := os.WriteFile("empty-files.txt", bs, 0644)
		if writeErr != nil {
			fmt.Printf("Error during writing file: %v", writeErr)
			return
		}
	}

}

func optimizedMemoryAllocation(files []os.DirEntry) int {
	var totalMem int
	for _, el := range files {
		info, _ := el.Info()
		if info.Size() == 0 {
			totalMem += len(info.Name()) + 1
		}
	}
	return totalMem
}
