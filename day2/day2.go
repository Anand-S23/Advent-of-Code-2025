package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func split_n(s string, n int) []string {
    if n <= 0 {
        return nil
    }

    partSize := (len(s) + n - 1) / n // ceiling division
    parts := []string{}

    for i := 0; i < len(s); i += partSize {
        end := i + partSize
        if end > len(s) {
            end = len(s)
        }
        parts = append(parts, s[i:end])
    }

    return parts
}

func is_invalid(id string) bool {
	for i := 2; i < len(id) + 1; i++ {
		split := split_n(id, i)
		all_equal := true
		for _, v := range split {
			//println(id, i, v, split[0])
			if v != split[0] {
				all_equal = false

			}
		}

		if all_equal {
			return true
		}
	}

	return false
}

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
			if is_invalid(current_id_str) {
				res += current_id
			}
		}
	}

	println(res)
}
