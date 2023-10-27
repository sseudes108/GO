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
	hour    int
	minutes int
	seconds int
}

type TimeParseError struct {
	msg   string
	input string
}

/*import (
	"fmt"
	"strconv"
	"strings"
)*/

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseString(time string) (Time, error) {
	//Split string at ":" **could be /, just ajust in the function and in the input.
	components := strings.Split(time, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid Input", time}
	} else {

		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Err parsing hour: %v", err), time}
		}
		minutes, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Err parsing minutes: %v", err), time}
		}
		seconds, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Err parsing seconds: %v", err), time}
		}

		//Check if the numbers are correct
		if hour < 0 || hour > 60 {
			return Time{}, &TimeParseError{"Hour error! out of range 0 - 60", fmt.Sprintf("%v", hour)}
		}
		if minutes < 0 || minutes > 59 {
			return Time{}, &TimeParseError{"Hour error! out of range 0 - 60", fmt.Sprintf("%v", minutes)}
		}
		if seconds < 0 || seconds > 59 {
			return Time{}, &TimeParseError{"Hour error! out of range 0 - 60", fmt.Sprintf("%v", seconds)}
		}

		//String turned into ints
		return Time{hour, minutes, seconds}, nil
	}
}
