//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// Example time string: 14:07:33
func ParseTime(input string) (Time, error) {
	//SPLIT THE STRING 14:07:33 PELO SEPARADOR DETERMINADO NA FUNÇÃO SPLIT ( : NO CASO)
	components := strings.Split(input, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{"Incorrect number of elements", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{"Error parsing hours: %v", err}
		}
	}

	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	} else { //ATOI CONVERTE CADA PARTE DA STRING SEPARADA ANTERIORMENTE PARA O VALOR EM INT
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Err parsing hour: %v", err), input}
		}
		minutes, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Err parsing minutes: %v", err), input}
		}
		seconds, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Err parsing seconds: %v", err), input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hour out of  range: 0 <= hour <= 23", fmt.Sprintf("%v", hour)}
		}

		if minutes > 59 || minutes < 0 {
			return Time{}, &TimeParseError{"minute out of  range: 0 <= minute <= 59", fmt.Sprintf("%v", minutes)}
		}

		if seconds > 59 || seconds < 0 {
			return Time{}, &TimeParseError{"seconds out of  range: 0 <= seconds <= 59", fmt.Sprintf("%v", seconds)}
		}

		return Time{hour, minutes, seconds}, nil
	}
}
