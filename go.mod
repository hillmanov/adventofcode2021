module main

go 1.17

replace adventofcode/utils => ./utils

require (
	adventofcode/utils v0.0.0
	github.com/olekukonko/tablewriter v0.0.5
)

require github.com/mattn/go-runewidth v0.0.9 // indirect
