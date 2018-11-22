package days

import (
	"encoding/csv"
//	"fmt"
	"io"
	"strconv"
)

func ReadSpreadsheet(fileReader io.Reader) [][]string {
	reader := csv.NewReader(fileReader)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return csvData
}

// For each row, determine the difference between the
// largest value and the smallest value; the checksum is the
// sum of all of these differences.
func Checksum(fileReader io.Reader) int {
	var (
		max int = 0
		min int = 0
    sum int = 0
  )

	csvData := ReadSpreadsheet(fileReader)

	for row := range csvData {
		max = 0
		min = 0
		for col := range csvData[row] {
			i, _ := strconv.Atoi(csvData[row][col])
			if ( i > max ) {
				max = i
			}
			if ( min == 0 || i < min ) {
				min = i
			}
			//fmt.Println(col,"::", i, "::", min, max)
		}
		sum = sum + max - min
		//fmt.Println( "--", sum)
	}
	return sum
}

// find the only two numbers in each row where one evenly
// divides the other - that is, where the result of the division
// operation is a whole number. They would like you to find those
// numbers on each line, divide them, and add up each line's result.
func divide(a int, b int) int {
	if ( a > b ) {
		if ( a % b == 0 ) {
			return a / b
		}
	} else if ( b > a) {
		if ( b % a == 0 ) {
			return b / a
		}
	}
	return 0
}

func testAll(v int, values []int) int {
//	fmt.Println("compare", v, values)
	for i := range values {
		r := divide(v, values[i])
		if ( r != 0 ) {
			return r
		}
	}
	return 0
}

func findDivisors(data []string) int {
	var values []int

	for idx := range data {
		i, _ := strconv.Atoi(data[idx])
		r := testAll(i, values)
		if ( r != 0 ) {
//			fmt.Println("findDivisors", data, r)
			return r
		}
		values = append(values, i)
	}
//	fmt.Println("findDivisors", data, values, 0)
	return 0
}

func DivisibleChecksum(fileReader io.Reader) int {
	var sum int = 0

	csvData := ReadSpreadsheet(fileReader)
	for row := range csvData {
		sum = sum + findDivisors(csvData[row])
	}
	return sum
}

