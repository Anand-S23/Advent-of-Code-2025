package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func to_int(x string) int {
	n, err := strconv.ParseInt(x, 10, 0)
	if err != nil {
		log.Fatal("could not convert", x, "to int")
	}

	return int(n)
}

func part1(input []string) int {
	res := 0

	for _, line := range input {
		max_joltage := -1
		for i := 0; i < (len(line) - 1); i++ {
			for j := i + 1; j < len(line); j++ {
				first := to_int(string(line[i]))
				second := to_int(string(line[j]))
				current_joltage := first * 10 + second
				if max_joltage < current_joltage {
					max_joltage = current_joltage
				}
			}
		}

		res += max_joltage
	}

	return res
}

func part2(input []string) int {
	res := 0

	for _, line := range input {
		current_place := 100000000000
		start_index := 0
		current_joltage := 0

		for i := 1; i < 13; i++ {
			max_battery := -1
			battery_index := -1
			for j := start_index; j < (len(line) - (12 - i)); j++ {
				current_battery := to_int(string(line[j]))
				if current_battery > max_battery {
					max_battery = current_battery
					battery_index = j
				}
			}

			start_index = battery_index + 1
			current_joltage += max_battery * current_place
			current_place /= 10
		}

		res += current_joltage
	}

	return res
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		current_line := scanner.Text()
		if current_line == "" {
			continue
		}
		lines = append(lines, current_line)
	}

	part1_res := part1(lines)
	println("Part 1 Result:", part1_res)

	part2_res := part2(lines)
	println("Part 2 Result:", part2_res)
}
