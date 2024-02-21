package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: pcap2mpeg input.pcap output.mpg")
		os.Exit(1)
	}

	pcapFile := os.Args[1]
	mpegFile := os.Args[2]

	handle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatalf("\x1b[31m%v\x1b[0m", err)
	}
	defer handle.Close()

	outputFile, err := os.Create(mpegFile)
	if err != nil {
		log.Fatalf("\x1b[31m%v\x1b[0m", err)
	}
	defer outputFile.Close()

	// Get the total number of packets in the pcap file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetCount := 0
	for range packetSource.Packets() {
		packetCount++
	}

	// Reset packet source to beginning
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatalf("\x1b[31m%v\x1b[0m", err)
	}
	packetSource = gopacket.NewPacketSource(handle, handle.LinkType())

	// Initialize progress variables
	const progressBarWidth = 50
	writeStartTime := time.Now()
	// Loop through each packet in the pcap file
	packetIndex := 0
	for packet := range packetSource.Packets() {
		packetIndex++
		// Find the UDP layer in the packet
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			// Write the UDP payload data to the output file
			_, err := outputFile.Write(udpLayer.LayerPayload())
			if err != nil {
				log.Fatalf("\x1b[31m%v\x1b[0m", err)
			}

			// Calculate progress
			progress := float64(packetIndex) / float64(packetCount)
			progressBar := "[" + progressBar(progress, progressBarWidth) + "]"
			fmt.Printf("\r\x1b[32mWriting %s %s %.2f%% (%s)\x1b[0m",
				mpegFile,
				progressBar,
				progress*100,
				time.Since(writeStartTime))

			// Flush output to update progress message
			os.Stdout.Sync()
		}
	}

	fmt.Println("\n👾\x1b[32m", mpegFile, "has been created successfully\x1b[0m 👾")
}

// progressBar returns a string representing a progress bar of given width and progress value
func progressBar(progress float64, width int) string {
	numBars := int(progress * float64(width))
	return fmt.Sprintf("%s%s", strings.Repeat("=", numBars), strings.Repeat(" ", width-numBars))
}
