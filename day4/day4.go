package main

import (
	"bufio"
	"log"
	"os"
)

func is_cell_valid(row, col, w, h int) bool {
	return !(col < 0 || row < 0 || col >= w || row >= h)
}

type cell struct {
	row int
	col int
}

func part2(input [][]rune) int {
	res := 0
	w := len(input[0])
	h := len(input)


	changes := true
	for changes {
		changes = false
		valid_cells := make([]cell, 0)

		for row := 0; row < h; row++ {
			for col := 0; col < w; col++ {
				if input[row][col] != '@' {
					continue
				}

				count := 0
				if is_cell_valid(row-1, col, w, h) && input[row-1][col] == '@' {
					count++
				}

				if is_cell_valid(row+1, col, w, h) && input[row+1][col] == '@' {
					count++
				}

				if is_cell_valid(row, col-1, w, h) && input[row][col-1] == '@' {
					count++
				}

				if is_cell_valid(row, col+1, w, h) && input[row][col+1] == '@' {
					count++
				}

				if is_cell_valid(row-1, col-1, w, h) && input[row-1][col-1] == '@' {
					count++
				}

				if is_cell_valid(row+1, col+1, w, h) && input[row+1][col+1] == '@' {
					count++
				}

				if is_cell_valid(row-1, col+1, w, h) && input[row-1][col+1] == '@' {
					count++
				}

				if is_cell_valid(row+1, col-1, w, h) && input[row+1][col-1] == '@' {
					count++
				}

				if count < 4 {
					valid_cells = append(valid_cells, cell{ row: row, col: col })
					changes = true
					res++
				}
			}

			for _, c := range valid_cells {
				input[c.row][c.col] = '.'
			}

		}
	}

	return res
}

func part1(input [][]rune) int {
	res := 0
	w := len(input[0])
	h := len(input)

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if input[row][col] != '@' {
				continue
			}

			count := 0
			if is_cell_valid(row-1, col, w, h) && input[row-1][col] == '@' {
				count++
			}

			if is_cell_valid(row+1, col, w, h) && input[row+1][col] == '@' {
				count++
			}

			if is_cell_valid(row, col-1, w, h) && input[row][col-1] == '@' {
				count++
			}

			if is_cell_valid(row, col+1, w, h) && input[row][col+1] == '@' {
				count++
			}

			if is_cell_valid(row-1, col-1, w, h) && input[row-1][col-1] == '@' {
				count++
			}

			if is_cell_valid(row+1, col+1, w, h) && input[row+1][col+1] == '@' {
				count++
			}

			if is_cell_valid(row-1, col+1, w, h) && input[row-1][col+1] == '@' {
				count++
			}

			if is_cell_valid(row+1, col-1, w, h) && input[row+1][col-1] == '@' {
				count++
			}

			if count < 4 {
				res++
			}
		}
	}

	return res
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := make([][]rune, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		current_line := scanner.Text()
		if current_line == "" {
			continue
		}

		runes := make([]rune, 0)
		for _, r := range current_line {
			runes = append(runes, r)
		}
		lines = append(lines, runes)
	}

	part1_res := part1(lines)
	println("Part 1 Result:", part1_res)

	part2_res := part2(lines)
	println("Part 2 Result:", part2_res)
} 
