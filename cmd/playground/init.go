package main

import (
	"encoding/csv"
	"fmt"
	base62Manager "github.com/aryahmph/url-shortener/common/hash/base62"
	counterModel "github.com/aryahmph/url-shortener/internal/model/counter"
	"os"
)

func main() {
	// Create a new CSV file
	file, err := os.Create("id-parameter.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write header
	header := []string{"id"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	hasher := base62Manager.NewBase62Manager()
	for i := 1; i < counterModel.CounterRange/10; i++ {
		row := []string{hasher.Hash(uint64(i))}
		// Write 10x times
		for j := 0; j < 10; j++ {
			err = writer.Write(row)
			if err != nil {
				fmt.Println("Error writing row:", err)
				return
			}
		}
	}

	// Flush any buffered data to the underlying writer (the file)
	writer.Flush()

	// Check for any error during the Flush operation
	err = writer.Error()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	fmt.Println("CSV file created successfully.")
}
