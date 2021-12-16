package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

type PacketExpression struct {
	Version           int
	TypeID            int
	Value             int
	PacketExpressions []PacketExpression
}

func Part1() Any {
	transmission := getInput()
	packet := hexToBin(transmission)

	_, expression := process(packet)

	return sumOfVersions(expression)
}

func Part2() Any {
	transmission := getInput()
	packet := hexToBin(transmission)

	_, expression := process(packet)

	return evaluate(expression)
}

func process(packet string) (string, PacketExpression) {
	pe := PacketExpression{}

	var version int
	var typeID string
	version, packet = binToInt(packet[0:3]), packet[3:]
	typeID, packet = packet[0:3], packet[3:]

	pe.Version = version
	pe.TypeID = binToInt(typeID)

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
		pe.Value = binToInt(encoded)
		return packet, pe

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
				var subPE PacketExpression
				subPacket, subPE = process(subPacket)
				pe.PacketExpressions = append(pe.PacketExpressions, subPE)
			}

			return packet, pe

		case "1": // Next 11 bits contain bin number that says how many sub packets there are
			var subPacketCount int
			subPacketCount, packet = binToInt(packet[:11]), packet[11:]

			for i := 0; i < int(subPacketCount); i++ {
				var subPE PacketExpression
				packet, subPE = process(packet)
				pe.PacketExpressions = append(pe.PacketExpressions, subPE)
			}

			return packet, pe
		}
	}

	return packet, pe
}

func sumOfVersions(pe PacketExpression) int {
	sum := pe.Version
	for _, subPE := range pe.PacketExpressions {
		sum += sumOfVersions(subPE)
	}
	return sum
}

func evaluate(pe PacketExpression) int {
	switch pe.TypeID {
	case 0:
		sum := 0
		for _, subPE := range pe.PacketExpressions {
			sum += evaluate(subPE)
		}
		return sum
	case 1:
		product := 1
		for _, subPE := range pe.PacketExpressions {
			product *= evaluate(subPE)
		}
		return product
	case 2:
		minValue := math.MaxInt
		for _, subPE := range pe.PacketExpressions {
			minValue = MinInt(minValue, evaluate(subPE))
		}
		return minValue
	case 3:
		maxValue := math.MinInt
		for _, subPE := range pe.PacketExpressions {
			maxValue = MaxInt(maxValue, evaluate(subPE))
		}
		return maxValue
	case 4:
		return pe.Value
	case 5:
		a, b := evaluate((pe.PacketExpressions)[0]), evaluate((pe.PacketExpressions)[1])
		if a > b {
			return 1
		} else {
			return 0
		}
	case 6:
		a, b := evaluate((pe.PacketExpressions)[0]), evaluate((pe.PacketExpressions)[1])
		if a < b {
			return 1
		} else {
			return 0
		}
	case 7:
		a, b := evaluate((pe.PacketExpressions)[0]), evaluate((pe.PacketExpressions)[1])
		if a == b {
			return 1
		} else {
			return 0
		}
	}
	return -1
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
