package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

// IGMPGroup represents information about an IGMP group.
type IGMPGroup struct {
	Group   string
	Version string
}

func getIGMPGroups(iface *net.Interface) ([]IGMPGroup, error) {
	file, err := os.Open("/proc/net/igmp")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var igmpGroups []IGMPGroup
	scanner := bufio.NewScanner(file)

	// Regular expression to match IGMP entries in the /proc/net/igmp file
	igmpRegex := regexp.MustCompile(fmt.Sprintf(`^\s*%d\s+%s\s+:\s+(\d+)\s+(V\d)\s*$`, iface.Index, iface.Name))

	for scanner.Scan() {
		line := scanner.Text()
		match := igmpRegex.FindStringSubmatch(line)
		if match != nil {
			group := IGMPGroup{
				Group:   strings.TrimSpace(match[1]),
				Version: match[2],
			}
			igmpGroups = append(igmpGroups, group)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return igmpGroups, nil
}

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

		igmpGroups, err := getIGMPGroups(&iface)
		if err != nil {
			fmt.Println("Error getting IGMP groups for interface", iface.Name, ":", err)
			continue
		}

		fmt.Println("  IGMP Groups:")
		for _, igmpGroup := range igmpGroups {
			fmt.Printf("    Group: %s, Version: %s\n", igmpGroup.Group, igmpGroup.Version)
		}

		fmt.Println()
	}
}

