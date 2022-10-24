package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// a)
	fmt.Println("\na)")
	const n_a = 12
	Z_a := makeMultTable(n_a - 1)
	printTable(Z_a, n_a)

	// b) Fra tabellen ser vi at tallene 1, 5, 7 og 11 har multiplikative inverser

	// c) Fra tabellen ser vi at tallene 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 har multiplikative inverser.
	fmt.Println("\nc)")
	const n_c = 11
	Z_c := makeMultTable(n_c - 1)
	printTable(Z_c, n_c)

	// d)
	fmt.Println("\nd)")
	const a_d = 11
	const n_d = 29

	mi := multInversBruteForce(a_d, n_d)
	if mi < 0 {
		fmt.Printf("Brute force could not find a multiplicative inverse for %d\n", a_d)
	} else {
		fmt.Printf("Brute force found multiplicative inverse %d, for %d (mod 29)\n", mi, a_d)
	}

	// e)
	// Det finnes én mult invers per tall a som ikke har noen felles faktorer med n

	// Eksempel:
	// n = 27. 27 er delelig med 3. Dermed har alle tall a (som er element i Z_27)
	// én mult invers, untatt de som er delelig på 3.

	// Dette betyr også at for primtall mengder, som 29, vil alle a (som er element i Z_29)
	// ha en mult invers.
}

func makeMultTable(n int) [][]int {
	Z := make([][]int, n)
	for i := 0; i < n; i++ {
		Z[i] = make([]int, n)
		for j := 0; j < n; j++ {
			Z[i][j] = (i + 1) * (j + 1)
		}
	}

	return Z
}

func makeRowString(nums []int) string {
	var res string
	for i := 0; i < len(nums); i++ {
		res += fmt.Sprintf("%4d", nums[i])
	}
	res += "\n"
	return res
}

func printTable(Z [][]int, n int) error {
	l := len(Z)
	if l != len(Z[0]) {
		return fmt.Errorf("2D array must be quadratic in shape")
	}

	fmt.Printf("Table: Z_%d\n", n)
	modCol := color.New(color.FgGreen).SprintFunc()
	regCol := color.New(color.FgRed).SprintFunc()

	fmt.Printf(strings.Repeat(" ", 5)+"|%s", makeRowString(Z[0]))
	fmt.Printf(strings.Repeat("-", 5) + "|" + strings.Repeat("----", l) + "\n")

	for i := 0; i < l; i++ {
		fmt.Printf("%4d |", Z[i][0])

		var row string
		for j := 0; j < l; j++ {
			var colFunc func(a ...interface{}) string
			if Z[i][j]%n == 1 {
				colFunc = modCol
			} else {
				colFunc = regCol
			}
			row += colFunc(fmt.Sprintf("%4d", Z[i][j]))
		}
		fmt.Println(row)
	}
	return nil
}

func multInversBruteForce(a, n int) int {
	for i := 0; i < n; i++ {
		if (a*i)%n == 1 {
			return i
		}
	}

	return -1
}
