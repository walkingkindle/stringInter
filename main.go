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

	// Remove duplicates
	arr = removeDuplicates(arr)

	// Sort the array
	sort.Ints(arr)

	// Output the sorted array
	fmt.Println("Output:", arr)
}

func removeDuplicates(arr []int) []int {
	encountered := make(map[int]bool)
	result := []int{}

	for _, num := range arr {
		if !encountered[num] {
			encountered[num] = true
			result = append(result, num)
		}
	}

	return result
}

