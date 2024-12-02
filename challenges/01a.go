package challenges

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day_01a() int {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/data/day-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int64, 0, 1000)
	right := make([]int64, 0, 1000)

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), "   ")
		left_int, _ := strconv.ParseInt(slice[0], 10, 64)
		right_int, _ := strconv.ParseInt(slice[1], 10, 64)
		left = append(left, left_int)
		right = append(right, right_int)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var sum int64 = 0
	for i := range left {
		// take absolute value of difference
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return int(sum)
}
