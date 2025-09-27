package encode // https://exercism.org/tracks/go/exercises/run-length-encoding

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	length := len(input)
	var output strings.Builder

	for left, right := 0, 0; left < length; left = right {

		for ; right < length && input[left] == input[right]; right++ {
			//
		}

		if right-left == 1 {
			output.WriteByte(input[left])
		} else {
			output.WriteString(strconv.Itoa(right - left))
			output.WriteByte(input[left])
		}
	}

	return output.String()
}

func RunLengthDecode(input string) string {
	length := len(input)
	var output strings.Builder

	for left, right := 0, 0; left < length; left = right {

		for ; right < length && input[right] >= '0' && input[right] <= '9'; right++ {
			//
		}

		if right-left != 0 {
			counter, err := strconv.Atoi(string(input[left:right]))
			if err != nil {
				return ""
			}

			output.WriteString(strings.Repeat(string(input[right]), counter))
		} else {
			output.WriteByte(input[right])
		}

		right++
	}

	return output.String()
}
