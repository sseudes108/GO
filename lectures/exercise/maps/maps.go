//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status displaying:
//   - number of servers
//   - number of servers for each status (Online, Offline, Maintenance, Retired)
func displayServerInfo(mapServers map[string]int) {
	println("Total of servers:", len(mapServers))

	//Instructor solution
	/*
		stats := make(map[int]int)

		for _, status := range mapServers {
			switch status {
			case Online:
				stats[Online]++
			case Offline:
				stats[Offline]++
			case Maintenance:
				stats[Maintenance]++
			case Retired:
				stats[Retired]++
			default:
				panic("Unhandled server status")
			}
		}

		println("Total of servers Online:", stats[Online])
		println("Total of servers Offline:", stats[Offline])
		println("Total of servers Under Maintence:", stats[Maintenance])
		println("Total of servers Retired:", stats[Retired])

	*/

	//Mine
	onlineServers := 0
	offlineServers := 0
	underMaintenceServers := 0
	retiredServers := 0

	for i, server := range mapServers {
		sr, found := mapServers[i]

		if found {
			if sr == Online {
				onlineServers++
			} else if sr == Offline {
				offlineServers++
			} else if sr == Maintenance {
				underMaintenceServers++
			} else {
				retiredServers++
			}
		} else {
			fmt.Println(server, "Server not found!")
			return
		}
	}
	println("Total of servers Online:", onlineServers)
	println("Total of servers Offline:", offlineServers)
	println("Total of servers Under Maintence:", underMaintenceServers)
	println("Total of servers Retired:", retiredServers)

}

// - change server status of all servers to `Maintenance`
func SetAllServersToMaintence(mapServers map[string]int) {
	for server := range mapServers {
		mapServers[server] = Maintenance
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	//* Create a map using the server names as the key and the server status
	//  as the value

	//Mine
	mapServers := make(map[string]int) /*{
		servers[0]: Online,
		servers[1]: Online,
		servers[2]: Online,
		servers[3]: Online,
		servers[4]: Online,
	}*/

	//instructor
	for _, server := range servers {
		mapServers[server] = Online
	}

	displayServerInfo(mapServers)

	//  - call display server info function
	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	//  - call display server info function
	//  - change server status of all servers to `Maintenance`
	//  - call display server info function
	println("Updating Servers")

	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	mapServers["darkstar"] = Retired
	mapServers["aiur"] = Offline

	//  - call display server info function
	displayServerInfo(mapServers)
	println("Apply Maintence")

	//  - change server status of all servers to `Maintenance`
	SetAllServersToMaintence(mapServers)

	println("Updated Status")

	//  - call display server info function
	displayServerInfo(mapServers)

}
