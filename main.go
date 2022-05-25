package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var line int
	for {
		_, err := fmt.Fscanf(fd, "%d\n", &line)
		if err != nil {

			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}
	return
}

func main() {
	nums := readFile("data.txt")

	// statistics need decimal numbers, hence will use float64 type
	sum := 0.0
	Average := 0.0
	Median := 0.0
	size := len(nums)

	// calculate the average

	for i := 0; i < size; i++ {
		sum += float64(nums[i])
		Average = sum / float64(size)

	}
	// this calculates the median

	sort.Ints(nums)
	M := size / 2

	// if size is odd, take the middle number
	if size%2 == 1 {
		Median = float64(nums[M])
	} else {
		// if size is even, take the average of the middle two numbers
		Median = (float64(nums[M-1]) + float64(nums[M])) / 2
	}

	sumSqErrors := 0.0
	for _, number := range nums {
		sumSqErrors += math.Pow(float64(number)-float64(Average), 2)
	}
	variance := sumSqErrors / float64(size)
	sd := math.Sqrt(variance)
	resultV := int(math.Round(variance))
	resultS := int(math.Round(sd))
	resultA := int(math.Round(Average))
	resultM := int(math.Round(Median))

	fmt.Println("Average:", resultA)
	fmt.Println("Median:", resultM)
	fmt.Println("Variance:", resultV)
	fmt.Println("Standard Deviation:", resultS)
}
