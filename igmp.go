package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		fmt.Println("Interface:", iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error getting addresses for interface", iface.Name, ":", err)
			continue
		}

		for _, addr := range addrs {
			fmt.Println("  Address:", addr)
		}

		groups, err := iface.MulticastAddrs()
		if err != nil {
			fmt.Println("Error getting multicast groups for interface", iface.Name, ":", err)
			continue
		}

		for _, group := range groups {
			fmt.Println("  Multicast Group:", group)
		}

		fmt.Println()
	}
}

