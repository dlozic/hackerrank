package algorithms

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	total := int64(0)
	min := int64(arr[0])
	max := int64(arr[0])

	for _, value := range arr {
		tmpVal := int64(value)
		total += tmpVal
		if tmpVal < min {
			min = tmpVal
		}
		if tmpVal > max {
			max = tmpVal
		}
	}

	minSum := total - min
	maxSum := total - max
	fmt.Printf("%d %d\n", maxSum, minSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}
