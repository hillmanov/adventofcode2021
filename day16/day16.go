package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	transmission := getInput()
	packet := hexToBin(transmission)

	var packetVersion int64
	process(packet, &packetVersion)

	return packetVersion
}

func Part2() Any {
	return nil
}

func process(packet string, packetVersion *int64) string {
	for len(packet) > 0 && strings.Contains(packet, "1") {
		var version int
		var typeID string
		version, packet = binToInt(packet[0:3]), packet[3:]
		typeID, packet = packet[0:3], packet[3:]

		*packetVersion += int64(version)

		switch typeID {

		case "100": // Literal value
			encoded := ""
			for {
				sentinal := string(packet[0])
				encoded, packet = encoded+packet[1:5], packet[5:]
				if sentinal == "0" {
					break
				}
			}
			// value, _ := strconv.ParseInt(encoded, 2, 64)
			packet = process(packet, packetVersion)
			return packet

		default: // Operator
			var lengthTypeID string
			lengthTypeID, packet = string(packet[0]), packet[1:]

			switch lengthTypeID {

			case "0": // Next 15 bits contain bin number that says how many bits the next sub packet(s) is/are
				var subPacketLength int
				var subPacket string

				subPacketLength, packet = binToInt(packet[:15]), packet[15:]
				subPacket, packet = packet[:subPacketLength], packet[subPacketLength:]

				for len(subPacket) > 0 {
					subPacket = process(subPacket, packetVersion)
				}

				return packet

			case "1": // Next 11 bits contain bin number that says how many sub packets there are
				var subPacketCount int
				subPacketCount, packet = binToInt(packet[:11]), packet[11:]

				for i := 0; i < int(subPacketCount); i++ {
					packet = process(packet, packetVersion)
				}

				return packet
			}
		}
	}

	return packet
}

func hexToBin(hex string) string {
	builder := []string{}
	for _, c := range strings.Trim(hex, "\n") {
		o, err := strconv.ParseInt(string(c), 16, 64)
		if err != nil {
			panic(err)
		}
		builder = append(builder, fmt.Sprintf("%04b", o))
	}
	return strings.Join(builder, "")
}

func binToInt(bin string) int {
	v, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func getInput() string {
	contents, _ := ReadContents(f, "input.txt")
	return contents
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 16: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 16: Part 2: = %+v\n", part2Solution)
}
