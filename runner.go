package main

import (
	. "adventofcode/utils"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)

	command := flag.String("command", "runCurrent", "Command to run")
	flag.Parse()

	switch *command {
	case "runCurrent":
		runCurrent()
	case "runAll":
		runAll()
	case "buildCurrent":
		buildCurrent()
	case "buildAll":
		buildAll()
	}
}

func buildCurrent() error {
	currentDay := getCurrentDay()
	log.Printf("Building %s", currentDay)

	return buildDay(currentDay)
}

func runCurrent() {
	if err := buildCurrent(); err != nil {
		log.Fatal(err)
	}

	currentDay := getCurrentDay()
	dayPluginPath := filepath.Join("./", currentDay, currentDay+".so")

	dayPlugin, _ := plugin.Open(dayPluginPath)

	part1Symbol, _ := dayPlugin.Lookup("Part1")
	part2Symbol, _ := dayPlugin.Lookup("Part2")

	part1 := part1Symbol.(func() Any)
	part2 := part2Symbol.(func() Any)

	log.Printf("Running %s", currentDay)
	part1Solution := part1()
	part2Solution := part2()

	fmt.Printf("part1Solution = %+v\n", part1Solution)
	fmt.Printf("part2Solution = %+v\n", part2Solution)
}

func buildAll() error {
	days := getDays()
	log.Printf("Building %d days", len(days))
	return nil
}

func runAll() {
	days := getDays()
	log.Printf("Running %d days", len(days))
}

func buildDay(day string) error {
	cmd := exec.Command("go", "build", "-buildmode", "plugin")
	cmd.Dir = "./" + day
	if _, err := cmd.Output(); err != nil {
		log.Printf("Error building %s: %s", day, err)
		return err
	}
	return nil
}

func runDay(day string) error {
	return nil
}

func getCurrentDay() string {
	days := getDays()
	return days[len(days)-1]
}

func getDayNum(day string) int {
	numString := strings.TrimPrefix(day, "day")
	num, _ := strconv.Atoi(numString)
	return num
}

func getDays() []string {
	days := []string{}
	fsEntries, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fsEntries {
		if strings.HasPrefix(f.Name(), "day") {
			days = append(days, f.Name())
		}
	}
	return days
}
