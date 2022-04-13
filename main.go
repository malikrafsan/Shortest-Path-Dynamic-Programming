package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MaxUint = ^uint(0)
var INF = int(MaxUint >> 1)

func describeMatrix(matrix [][]int) {
	// debugging purpose

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == INF {
				fmt.Printf("-1 ")
			} else {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

func solve(matrix [][]int) {
	dist := make([]int, len(matrix))
	path := make([]int, len(matrix)-1)

	dist[len(matrix)-1] = 0

	for i := len(matrix) - 2; i >= 0; i-- {
		// inisialisasi dengan infinity
		dist[i] = INF
		temp := i

		// cari node terminimum yang mengarah ke node ke-i
		for j := i; j < len(matrix); j++ {
			if matrix[i][j] != INF { // jika ada jalan
				if (dist[j] + matrix[i][j]) < dist[i] {
					dist[i] = dist[j] + matrix[i][j]
					temp = j
				}
			}
		}
		// simpan node sebelumnya dalam path
		path[i] = temp + 1
	}

	final := make([]int, 0)
	final = append(final, 1)

	// cari path yang mengarah ke node tujuan
	iterator := 0
	for true {
		final = append(final, path[iterator])
		if iterator >= len(path) || path[iterator] == len(matrix) {
			break
		}

		iterator = path[iterator] - 1
	}

	// tampilkan hasil ke layar
	fmt.Print("Path: ", 1)
	for i := 1; i < len(final); i++ {
		fmt.Printf(" -> %d", final[i])
	}
	fmt.Println("\ndistance:", dist[0])
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	// baca matrix input dari file
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// parse input matrix
	sliceData := strings.Split(string(data), "\n")
	matrix := make([][]int, len(sliceData))

	for i := 0; i < len(sliceData); i++ {
		splitted := strings.Split(strings.TrimRight(sliceData[i], "\r\n"), " ")
		temp := make([]int, len(splitted))
		for j := 0; j < len(splitted); j++ {
			converted, err := strconv.Atoi(splitted[j])
			if err != nil {
				panic(err)
			}

			if converted == -1 {
				temp[j] = INF
			} else {
				temp[j] = converted
			}
		}

		for j := 0; j < len(temp); j++ {
			matrix[i] = temp
		}
	}

	// cari penyelesaian dari matrix input
	solve(matrix)
}
