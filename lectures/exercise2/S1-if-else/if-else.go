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

func isWeekDay(day int) bool {
	if day > friday {
		return false
	} else {
		return true
	}
}

func returnMessage(key int, weekday int) {
	if key == 0 {
		fmt.Println("Acess Denied", weekday)
	}
	if key == 1 {
		fmt.Println("Acess Granted", weekday)
	}
}

func checkAcess(user int, day int) {
	if user == admin || user == manager {
		returnMessage(1, day)
	}

	if user == contractor {
		if isWeekDay(day) {
			returnMessage(0, day)
		} else {
			returnMessage(1, day)
		}
	}

	if user == member {
		if isWeekDay(day) {
			returnMessage(1, day)
		} else {
			returnMessage(0, day)
		}
	}

	if user == guest {
		if day == monday || day == wednesday || day == friday {
			returnMessage(1, day)
		} else {
			returnMessage(0, day)
		}
	}
}

func main() {
	fmt.Println("Start")

	//   - Access at any time: Admin, Manager
	fmt.Printf("Admin: ")
	checkAcess(admin, friday) //granted 4

	//   - Access weekdays: Member
	fmt.Printf("Manager: ")
	checkAcess(manager, monday) //granted 0

	//   - Access weekends: Contractor
	fmt.Printf("Contractor: ")
	checkAcess(contractor, monday) //denied 2

	//   - Access weekdays: Member
	fmt.Printf("Member: ")
	checkAcess(member, wednesday) //granted 2

	//   - Access Mondays, Wednesdays, and Fridays: Guest
	fmt.Printf("Guest: ")
	checkAcess(guest, thursday) // denied 3

	fmt.Println("End")
}
