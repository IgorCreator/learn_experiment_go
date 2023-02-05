package main

import (
	"fmt"
	"os"
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
	emptyFiles := make([]byte, 0, totalMem)

	for _, el := range files {
		info, _ := el.Info()

		if info.Size() == 0 {
			emptyFiles = append(emptyFiles, info.Name()...)
			emptyFiles = append(emptyFiles, "\n"...)
			continue
		}

		if info.IsDir() {
			fmt.Printf("Inner Dir name: \"%s\". Size: %d Bytes \n", info.Name(), info.Size())
		}

		fmt.Printf("File name: \"%s\". Size: %d Bytes \n", info.Name(), info.Size())

	}

	if len(emptyFiles) != 0 {
		fmt.Printf("\nEmpty files: %s", emptyFiles)
		writeErr := os.WriteFile("empty-files.txt", emptyFiles, 0644)
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
