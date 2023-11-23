// --Summary:
//
//	The existing program is used to restrict access to a resource
//	based on day of the week and user role. Currently, the program
//	allows any user to access the resource. Implement the functionality
//	needed to grant resource access using any combination of `if`, `else if`,
//	and `else`.
//
// --Requirements:
//   - Use the accessGranted() and accessDenied() functions to display
//     informational messages
//   - Access at any time: Admin, Manager
//   - Access weekends: Contractor
//   - Access weekdays: Member
//   - Access Mondays, Wednesdays, and Fridays: Guest
package main

import "fmt"

const (
	monday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	admin = iota
	manager
	contractor
	member
	guest
)

func accessGranted() {
	fmt.Println("Acess Granted")
}

func accessDenied() {
	fmt.Println("Acess Denied")
}

func isWeekday(weekday int) bool {
	if weekday <= 4 {
		return true
	} else {
		return false
	}
}

func checkAcess(role int, day int) {
	weekday := isWeekday(day)

	if role == admin || role == manager {
		accessGranted()
	} else if role == contractor && !weekday {
		accessGranted()
	} else if role == member && weekday {
		accessGranted()
	} else if role == guest {
		if day == monday || day == wednesday || day == friday {
			accessGranted()
		}
	} else {
		accessDenied()
	}
}

func main() {

	// --Requirements:
	//   - Use the accessGranted() and accessDenied() functions to display
	//     informational messages
	//   - Access at any time: Admin, Manager
	//   - Access weekends: Contractor
	//   - Access weekdays: Member
	//   - Access Mondays, Wednesdays, and Fridays: Guest
	fmt.Print("Start\n")

	checkAcess(admin, sunday)         //granted
	checkAcess(manager, friday)       //granted
	checkAcess(contractor, wednesday) //denied
	checkAcess(member, saturday)      //denied
	checkAcess(guest, wednesday)      //granted

	fmt.Print("End\n")
}
