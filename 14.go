package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readmatrix("14_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
	solve(data)
}

func readmatrix(filename string) ([][]int, error) {
	var result [][]int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for i := 0; scanner.Scan(); i++ {
		line = scanner.Text()
		result = append(result, []int{})
		for _, s := range strings.Split(line, " ") {
			value, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			result[i] = append(result[i], value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func solve(graph [][]int) {
	// var visited []int

	points := make([]float64, len(graph))
	for i := range len(graph)-1 {
		points[i+1] = math.Inf(1)
	}

	var startpoint int = 0

	fmt.Println(neighbors(graph[startpoint]))
}

func neighbors(connections []int) []int {

	var neighbors []int 

	for i, v := range connections {
		if v != 0 {
			neighbors = append(neighbors, i)
		}
	}

	return neighbors
}
