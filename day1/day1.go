package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"fmt"
)

func mod(a, b int) int {
    r := a % b
    if r < 0 { r += b }
    return r
}

func update_tick(text string, current_tick int) int {
	dir := text[0]
	distance, err := strconv.Atoi(text[1:])
    if err != nil {
		log.Fatal(err)
    }

	if dir == 'R' {
		current_tick += distance
	} else if dir == 'L' {
		current_tick -= distance
	} else {
		log.Fatal("Not right or left")
	}

	current_tick = mod(current_tick, 100)
	fmt.Printf("%c, %d\n", dir, current_tick)
	return current_tick
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
		current_tick = update_tick(scanner.Text(), current_tick)
		if current_tick == 0 {
			res++
		}
	}

	println(res)
}
