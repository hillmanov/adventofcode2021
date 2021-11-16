define GO_MOD_TEMPLATE
module adventofcode/day${day}

go 1.17

replace adventofcode/utils => ../utils

require adventofcode/utils v0.0.0
endef

define GO_FILE_TEMPLATE
package main

import (
	"adventofcode/utils"
	"fmt"
)

func Part1() *int {
	return nil
}

func Part2() *int {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

  fmt.Printf("Day ${day}: Part 1: = %+v\\n", part1Solution)
	fmt.Printf("Day ${day}: Part 2: = %+v\\n", part2Solution)
}
endef

export GO_MOD_TEMPLATE
export GO_FILE_TEMPLATE

init:
	@mkdir day${day}
	@echo "$$GO_MOD_TEMPLATE" > day${day}/go.mod
	@echo "$$GO_FILE_TEMPLATE" > day${day}/day${day}.go
	@touch day${day}/input.txt
	@touch day${day}/README.md

build-all:

time-all:

time:
