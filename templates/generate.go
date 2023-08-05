package main

import (
	"fmt"
	"os"
)

func main() {
	dirs := []string{
		"internal",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	fmt.Println("Golang project structure generated successfully!")
}
