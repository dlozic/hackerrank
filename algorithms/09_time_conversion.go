package algorithms

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	var hour, min, sec int
	var period string

	fmt.Sscanf(s, "%d:%d:%d%s", &hour, &min, &sec, &period)

	if period == "PM" && hour != 12 {
		hour += 12
	} else if period == "AM" && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}
