package encode // https://exercism.org/tracks/go/exercises/run-length-encoding

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	length := len(input)
	output := make([]string, 0, length)

	for left, right := 0, 0; left < length; left = right {

		for ; right < length && input[left] == input[right]; right++ {
			//
		}

		if right-left == 1 {
			output = append(output, string(input[left]))
		} else {
			output = append(output, strconv.Itoa(right-left)+string(input[left]))
		}
	}

	return strings.Join(output, "")
}

func RunLengthDecode(input string) string {
	var output []string
	length := len(input)

	for left, right := 0, 0; left < length; left = right {

		for ; right < length && input[right] >= '0' && input[right] <= '9'; right++ {
			//
		}

		if right-left != 0 {
			counter, err := strconv.Atoi(string(input[left:right]))
			if err != nil {
				return ""
			}

			output = append(output, strings.Repeat(string(input[right]), counter))
		} else {
			output = append(output, string(input[right]))
		}

		right++
	}

	return strings.Join(output, "")
}
