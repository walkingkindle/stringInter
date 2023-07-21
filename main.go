package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"
)

func main() {
	inputStr := flag.String("input", "", "Input string representing an array")
	flag.Parse()

	if *inputStr == "" {
		fmt.Println("Error: Input string is empty. Please provide a valid input.")
		return
	}

	trimmedInput := strings.TrimSpace(*inputStr)
	if !strings.HasPrefix(trimmedInput, "[") || !strings.HasSuffix(trimmedInput, "]") {
		fmt.Println("Error: Invalid input format. The input should start with '[' and end with ']'.")
		return
	}

	var arr []int
	err := json.Unmarshal([]byte(trimmedInput), &arr)
	if err != nil {
		fmt.Println("Error: Invalid input format. The input should be a valid array of integers.")
		return
	}

	arr = deduplicateInPlace(arr)

	sort.Ints(arr)

	fmt.Println("Output:", arr)
}

func deduplicateInPlace(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	sort.Ints(arr)

	uniqueIdx := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[uniqueIdx] {
			uniqueIdx++
			arr[uniqueIdx] = arr[i]
		}
	}

	return arr[:uniqueIdx+1]
}

