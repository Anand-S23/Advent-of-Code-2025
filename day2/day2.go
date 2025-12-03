package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := ""
	for scanner.Scan() {
		input = scanner.Text()
		break
	}

	res := 0 

	code_set := strings.Split(input, ",")
	for _, v := range code_set {
		codes := strings.Split(v, "-")
		// println(v)

		start, err := strconv.Atoi(codes[0])
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(codes[1])
		if err != nil {
			log.Fatal(err)
		}

		for current_id := start; current_id < end + 1; current_id++ {
			current_id_str := strconv.Itoa(current_id)
			if len(current_id_str) % 2 == 0 {
				half := len(current_id_str) / 2
				if current_id_str[:half] == current_id_str[half:] {
					res += current_id
				}
			}
		}
	}

	println(res)
}
