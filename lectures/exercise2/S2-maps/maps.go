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
//  - call display server info function-
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

type Server string

func changeServerStatus(servMap map[Server]int, server Server, status int) {
	fmt.Println("Changing server", server, "status")
	servMap[server] = status
}

func maintenceAllServers(servMap map[Server]int) {
	fmt.Println("Setting Maintence")
	for server := range servMap {
		servMap[server] = Maintenance
	}
}

func OverallServerStatus(servMap map[Server]int) {

	fmt.Println("Total servers: ", len(servMap))

	online, offline, maintenance, retired := 0, 0, 0, 0

	for server, status := range servMap {
		switch status {
		case Online:
			online++
		case Offline:
			offline++
		case Maintenance:
			maintenance++
		case Retired:
			retired++
		}
		fmt.Println("Server: ", server, "- Status: ", statusToText(status))
	}
}

func statusToText(number int) string {
	if number == Online {
		return "Online"
	} else if number == Offline {
		return "Offline"
	} else if number == Maintenance {
		return "Maintenance"
	} else if number == Retired {
		return "Retired"
	}
	return "Error"
}

func main() {
	fmt.Println("Start")

	servers := []Server{"darkstar", "aiur", "omicron", "w359", "baseline"}
	mapServers := make(map[Server]int)

	for _, server := range servers {
		mapServers[server] = Online
	}

	//  - call display server info function
	OverallServerStatus(mapServers)

	//  - change server status of `darkstar` to `Retired`
	changeServerStatus(mapServers, "darkstar", Retired)
	//  - change server status of `aiur` to `Offline`
	changeServerStatus(mapServers, "aiur", Offline)

	//  - call display server info function
	OverallServerStatus(mapServers)

	//  - change server status of all servers to `Maintenance`
	maintenceAllServers(mapServers)

	//  - call display server info function
	OverallServerStatus(mapServers)

	fmt.Println("End")
}
