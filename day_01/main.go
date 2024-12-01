package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("---- day 1 -----")
	if len(os.Args) < 2 {
		fmt.Println("Input is not provided")
		return
	}
	args := os.Args[1:]
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Failed to read input ", args[0])
		return
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("invalid line format")
		} else {
			left_int, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("invalid left number", err)
			} else {
				left = append(left, left_int)
			}
			right_int, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("invalid right number")
			} else {
				right = append(right, right_int)
			}
		}
	}
	fmt.Println("finished parsing")
	if len(left) != len(right) {
		fmt.Println("Both lists do not have the same len")
		return
	}
	sort.Ints(left)
	sort.Ints(right)
	var distance int64
	for i := range left {
		distance += int64(math.Abs((float64)(right[i] - left[i])))
	}
	fmt.Println("distance is ", distance)

	// part 2
	// can use hash map or can use the sorted structure
	var score int
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			l++
		} else if left[l] > right[r] {
			r++
		} else { // equal numbers
			var count int
			val := right[r]
			for r < len(right) && val == right[r] {
				count++
				r++
			}
			for l < len(left) && left[l] == val {
				score += count * val
				l++
			}
		}
	}
	fmt.Println("smililarity score is ", score)
}
