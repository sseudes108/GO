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

type Server string

const (
	online = iota
	offline
	maintenance
	retired
)

func showServersStatus(servers map[Server]int) {
	fmt.Println("Number os servers:", len(servers))
	stats := ""
	onlines, offlines, uMaintence, retireds := 0, 0, 0, 0
	for server, status := range servers {

		switch status {
		case online:
			stats = "Online"
			onlines++
		case offline:
			stats = "Offline"
			offlines++
		case maintenance:
			stats = "Maintence"
			uMaintence++
		case retired:
			stats = "Retired"
			retireds++
		}
		fmt.Println("Server:", server, "-", stats)
	}
	fmt.Println("----")
	fmt.Println("Servers Online", onlines)
	fmt.Println("Servers Offline", offlines)
	fmt.Println("Servers Under Maintence", uMaintence)
	fmt.Println("Servers Retired", retireds)
	fmt.Println("----")
}

func setMaintence(servers map[Server]int) {
	for server := range servers {
		servers[server] = maintenance
	}
}

func main() {
	servers := []Server{"darkstar", "aiur", "omicron", "w359", "baseline"}

	mapsServer := make(map[Server]int)

	for _, server := range servers {
		mapsServer[server] = online
	}

	showServersStatus(mapsServer)

	mapsServer["darkstar"] = retired
	mapsServer["aiur"] = offline

	showServersStatus(mapsServer)

	setMaintence(mapsServer)

	showServersStatus(mapsServer)
}
