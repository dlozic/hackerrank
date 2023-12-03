package algorithms

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var pos, neg, zero int32 = 0, 0, 0

	for _, num := range arr {
		switch {
		case num > 0:
			pos++
		case num < 0:
			neg++
		case num == 0:
			zero++
		}
	}

	count := len(arr)

	fmt.Printf("%.6f\n", float64(pos)/float64(count))
	fmt.Printf("%.6f\n", float64(neg)/float64(count))
	fmt.Printf("%.6f\n", float64(zero)/float64(count))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}
