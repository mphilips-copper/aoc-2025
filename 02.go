package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("02.txt")
	if err != nil {
		log.Fatal(err)
	}

	stringRanges := strings.Split(string(input), ",")
	sumInvalidIds := 0

	for _, stringRange := range stringRanges {

		stringIds := strings.Split(stringRange, "-")
		// let's assume input is perfect and there will be exactly 2 ids here
		start, _ := strconv.Atoi(stringIds[0])
		end, _ := strconv.Atoi(stringIds[1])

		ids := make(map[int]string)
		for ; start <= end; start++ {
			ids[start] = strconv.Itoa(start)
		}

		for key, value := range ids {
			// ids with an odd number of digits are valid for sure
			if len(value) % 2 != 0 {
				continue
			}

			// from google ai overview
			mid := len(value) / 2
			firstHalf := value[:mid]
			secondHalf := value[mid:]
			//

			if firstHalf == secondHalf {
				sumInvalidIds = sumInvalidIds + key
			}
		}
	}

	fmt.Println(sumInvalidIds)
}
