//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func isWeekDay(today int) bool {
	if today <= Friday {
		return true
	} else {
		return false
	}
}

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	// The day and role. Change these to check your work.
	var today, role = Saturday, Member

	if role == Admin || role == Manager { //* Access at any time: Admin, Manager
		accessGranted()
		return
	} else if role == Contractor && !isWeekDay(today) { //* Access weekends: Contractor
		accessGranted()
		return
	} else if role == Member && isWeekDay(today) { //* Access weekdays: Member
		accessGranted()
		return
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) { //* Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
		return
	} else {
		accessDenied()
		return
	}
}
