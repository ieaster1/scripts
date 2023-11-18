package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ouis = map[string][]int{
		"xen":  {0x00, 0x16, 0x3E},
		"qemu": {0x52, 0x54, 0x00},
	}
)

var (
	countFlag int
	ouiFlag   string
)

func init() {
	flag.IntVar(&countFlag, "c", 1, "number of MACs to create")
	flag.StringVar(&ouiFlag, "o", "", "input for OUI, e.g., \"00:12:ac\"")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-h] [-c] [-o]\n\nCreate random MAC with Locally Administered Organizational Unique Identifier.\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "optional arguments:\n")
		flag.PrintDefaults()
	}
}

func createRandomMAC() ([]int, string) {
	var oui []int

	if ouiFlag != "" {
		strOui := strings.Split(ouiFlag, ":")
		for _, hex := range strOui {
			i, _ := strconv.ParseInt(hex, 16, 64)
			oui = append(oui, int(i))
		}
	} else {
		var ok bool
		oui, ok = ouis["qemu"]
		if !ok {
			oui = ouis["xen"]
		}
	}

	counter := rand.Int63n(1 << 24) // generate a random 24-bit counter
	decimalMAC := append(oui, int((counter>>16)&0xFF), int((counter>>8)&0xFF), int(counter&0xFF))
	mac := formatMAC(decimalMAC)
	return decimalMAC, mac
}

func formatMAC(decimalMAC []int) string {
	var hexMAC []string
	for _, octet := range decimalMAC {
		hexMAC = append(hexMAC, fmt.Sprintf("%02x", octet))
	}
	return strings.Join(hexMAC, ":")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	if ouiFlag != "" {
		// Make the regular expression case-insensitive
		if ok, _ := regexp.MatchString("[0-9a-f]{2}([:]?[0-9a-f]{2}){2}$", strings.ToLower(ouiFlag)); !ok {
			fmt.Printf("%s is incorrect format, check help for details\n", ouiFlag)
			os.Exit(1)
		}
	}

	for i := 0; i < countFlag; i++ {
		_, mac := createRandomMAC()
		fmt.Println(mac)
	}
}

