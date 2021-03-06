package utils

import (
	"bufio"
	"embed"
	"strconv"
)

type Any interface{}

func ReadLines(f embed.FS, filename string) ([]string, error) {
	file, err := f.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return s, nil
}

func ReadInts(f embed.FS, filename string) ([]int, error) {
	file, err := f.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	i := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		i = append(i, n)
	}

	return i, nil
}

func ReadContents(f embed.FS, filename string) (string, error) {
	contents, err := f.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func ReplaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func ParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinOf(numbers []int) int {
	var min int = numbers[0]
	for _, value := range numbers {
		if min > value {
			min = value
		}
	}
	return min
}

func MaxOf(numbers []int) int {
	var max int = numbers[0]
	for _, value := range numbers {
		if max < value {
			max = value
		}
	}
	return max
}

func MinMax(numbers []int) (int, int) {
	var max int = numbers[0]
	var min int = numbers[0]
	for _, value := range numbers {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func SumOf(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func UniqueOfString(strings []string) []string {
	keys := make(map[string]bool)
	unique := []string{}
	for _, entry := range strings {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique = append(unique, entry)
		}
	}
	return unique
}

func UniqueOfInt(numbers []int) []int {
	keys := make(map[int]bool)
	unique := []int{}
	for _, entry := range numbers {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique = append(unique, entry)
		}
	}
	return unique
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n

}

func CopyOf(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func IndexOf(haystack []int, needle int) int {
	for index, value := range haystack {
		if value == needle {
			return index
		}
	}
	return -1
}
