package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func mod(a, b int) int {
    r := a % b
    if r < 0 { r += b }
    return r
}

func update_tick(text string, current_tick int) (int, int)  {
	last_tick := current_tick
	dir := text[0]
	distance, err := strconv.Atoi(text[1:])
    if err != nil {
		log.Fatal(err)
    }

	if dir == 'R' {
		current_tick += distance % 100
	} else if dir == 'L' {
		current_tick -= distance % 100
	} else {
		log.Fatal("Not right or left")
	}

	rotations := (distance - mod(distance, 100)) / 100

	if current_tick < 0 || current_tick > 99 {
		if last_tick != 0 {
			rotations++
		}
	}

	if current_tick == 0 {
		rotations++
	} 

	current_tick = mod(current_tick, 100)
	return current_tick, rotations
}

func main() {
	current_tick := 50

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := 0
	for scanner.Scan() {
		new_tick, rotations := update_tick(scanner.Text(), current_tick)
		res += rotations
		current_tick = new_tick
	}

	println(res)
}
