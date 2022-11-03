package main

import (
	"fmt"
	"strings"
)

const ALPH_HEX = "0123456789abcdef"

var ALPH_RUNES = []rune(ALPH_HEX)
var RUNE_MAP, INT_MAP = createAlphabetMap()
var ALPH_LEN = len(ALPH_HEX)

func createAlphabetMap() (map[rune]int, map[int]rune) {
	runeMap := make(map[rune]int)
	byteMap := make(map[int]rune)

	for i, symbol := range ALPH_RUNES {
		runeMap[symbol] = i
		byteMap[i] = symbol
	}
	return runeMap, byteMap
}

func splitHexStringToInts(s string) (row int, col int) {
	runes := []rune(s)
	return RUNE_MAP[runes[0]], RUNE_MAP[runes[1]]
}

func tableToString(t [][]StateValue) string {
	var sb strings.Builder

	for _, row := range t {
		sb.WriteString(" ")
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%02x ", val))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func initEmptyState(size int) [][]StateValue {
	res := make([][]StateValue, size)
	for i := 0; i < size; i++ {
		res[i] = make([]StateValue, size)
	}
	return res
}

func padString(s string, l int) string {
	for i := len(s); i < l; i++ {
		s += " "
	}
	return s
}

func splitToBlocks(s string, l int) []string {
	sLen := len(s)
	iter := sLen / l

	if iter == 0 {
		return []string{padString(s, l)}
	}
	var blocks []string
	for i := 0; i < iter; i++ {
		blocks = append(blocks, s[i*l:(i+1)*l])
	}

	if sLen%l != 0 {
		blocks = append(blocks, padString(s[iter*l:], l))
	}

	return blocks
}

func textToStateBlocks(msg string, size int) [][][]StateValue {
	blocks := splitToBlocks(msg, size*size)

	states := make([][][]StateValue, len(blocks))

	for i, block := range blocks {
		states[i] = initEmptyState(size)
		values := []StateValue(block)
		valIndex := 0
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				states[i][k][j] = values[valIndex]
				valIndex++
			}
		}
	}

	return states
}

func stateBlocksToText(states [][][]StateValue) string {
	var sb strings.Builder
	for _, block := range states {
		for j, row := range block {
			for k := range row {
				sb.WriteByte(byte(block[k][j]))
			}
		}
	}

	return sb.String()
}

func xor(a, b StateValue) StateValue {
	return a ^ b
}

func xorTable(state, other [][]StateValue) [][]StateValue {
	numRows := len(state)
	numCols := len(state[0])
	newState := make([][]StateValue, numRows)

	for i, row := range state {
		newState[i] = make([]StateValue, numCols)
		for j, val := range row {
			newState[i][j] = xor(val, other[i][j])
		}
	}

	return newState
}

func getTableColumn(t [][]StateValue, colNum int) [][]StateValue {
	column := make([][]StateValue, len(t))

	for i := 0; i < len(column); i++ {
		column[i] = make([]StateValue, 1)
		column[i][0] = t[i][colNum]
	}

	return column
}

func appendTableColumn(t [][]StateValue, column [][]StateValue) [][]StateValue {
	if len(t[0]) == 0 {
		return column
	}

	numRows := len(t)
	numCols := len(t[0]) + 1

	newTable := make([][]StateValue, numRows)
	for i := 0; i < numRows; i++ {
		newTable[i] = make([]StateValue, numCols)

		for j := 0; j < numCols-1; j++ {
			newTable[i][j] = t[i][j]
		}
	}

	for i := 0; i < numRows; i++ {
		newTable[i][numCols-1] = column[i][0]
	}

	return newTable
}
