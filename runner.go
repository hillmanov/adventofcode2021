package main

import (
	. "adventofcode/utils"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type RunResult struct {
	Day       int
	Part      int
	Solution  Any
	StartTime time.Time
	EndTime   time.Time
}

// get the total time taken to run the solution
func (r RunResult) elapsedTime() int64 {
	return r.EndTime.Sub(r.StartTime).Milliseconds()
}

func (r RunResult) tableData() []string {
	return []string{
		strconv.Itoa(r.Day),
		strconv.Itoa(r.Part),
		fmt.Sprintf("%+v", r.Solution),
		fmt.Sprintf("%d", r.elapsedTime()),
	}
}

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
	results, _ := runDay(currentDay)
	renderResults(results)
}

func buildAll() error {
	days := getDays()
	for _, day := range days {
		if err := buildDay(day); err != nil {
			log.Fatalf("Error building %s: %s", day, err)
			return err
		}
	}
	return nil
}

func runAll() {
	if err := buildAll(); err != nil {
		log.Fatal(err)
	}

	runResults := []RunResult{}
	days := getDays()
	for _, day := range days {
		results, _ := runDay(day)
		runResults = append(runResults, results...)
	}
	renderResults(runResults)
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

func runDay(day string) ([]RunResult, error) {
	dayPluginPath := filepath.Join("./", day, day+".so")

	dayPlugin, _ := plugin.Open(dayPluginPath)

	part1Symbol, _ := dayPlugin.Lookup("Part1")
	part2Symbol, _ := dayPlugin.Lookup("Part2")

	part1 := part1Symbol.(func() Any)
	part2 := part2Symbol.(func() Any)

	log.Printf("Running %s", day)

	part1Result := RunResult{
		Day:       getDayNum(day),
		Part:      1,
		StartTime: time.Now(),
	}
	part1Result.Solution = part1()
	part1Result.EndTime = time.Now()

	part2Result := RunResult{
		Day:       getDayNum(day),
		Part:      2,
		StartTime: time.Now(),
	}
	part2Result.Solution = part2()
	part2Result.EndTime = time.Now()

	return []RunResult{part1Result, part2Result}, nil
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

func renderResults(rs []RunResult) {
	totalRunTime := int64(0)
	for _, r := range rs {
		totalRunTime += r.elapsedTime()
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Part", "Solution", "Time"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.SetFooter([]string{"", "", "Total", strconv.Itoa(int(totalRunTime))})
	for _, v := range rs {
		table.Append(v.tableData())
	}
	table.Render()
}
