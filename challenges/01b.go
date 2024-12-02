package challenges

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day_01b() int {
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

	num_of_occurances := make([]int64, 1000)
	for i := range left {
		for j := range right {
			if left[i] == right[j] {
				num_of_occurances[i]++
			}
		}
	}

	var sum int64 = 0
	for i := range left {
		sum += left[i] * num_of_occurances[i]
	}

	return int(sum)
}
