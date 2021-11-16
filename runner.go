package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range entries {
		if strings.HasPrefix(f.Name(), "day") {
			fmt.Println(f.Name())
		}
	}
}
