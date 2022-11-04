package main

import (
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

// Simple and probably crackable padding, due to time constraints
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

func textToStateBlocks(msg string, size int) []State {
	blocks := splitToBlocks(msg, size*size)

	states := make([]State, len(blocks))

	for i, block := range blocks {
		states[i] = NewZeroedState(size, size)
		values := []byte(block)
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

func stateBlocksToText(states []State) string {
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
