package main

import (
	"fmt"
	"strings"
)

type StateValue byte

var SUB_BYTES_TABLE = [][]StateValue{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
}

// var R_CONSTANTS = [][][]StateValue{
// 	{
// 		{0x00, 0x00},
// 		{0x01, 0x00},
// 	},
// 	{
// 		{0x00, 0x00},
// 		{0x02, 0x00},
// 	},
// 	{
// 		{0x00, 0x00},
// 		{0x04, 0x00},
// 	},
// }

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

func splitHexStringToInts(s string) (row int, col int) {
	runes := []rune(s)
	return RUNE_MAP[runes[0]], RUNE_MAP[runes[1]]
}

func subByte(val StateValue) StateValue {
	row, col := splitHexStringToInts(fmt.Sprintf("%02x", val))
	return SUB_BYTES_TABLE[row][col]
}

func subBytes(state [][]StateValue) [][]StateValue {
	for i, row := range state {
		for j, val := range row {
			state[i][j] = subByte(val)
		}
	}

	return state
}

func shiftRows(state [][]StateValue) [][]StateValue {
	maxShift := len(state)
	newState := initEmptyState(maxShift)

	for i, row := range state {
		for j := range row {
			newState[i][j] = state[i][(i+j)%maxShift]
		}
	}

	return newState
}

func f(a StateValue) StateValue {
	var c StateValue
	b := a << 1

	if a >= 0x80 {
		c = 0x1B
	} else {
		c = 0x00
	}

	return a ^ b ^ c
}

func mixColumns(state [][]StateValue) [][]StateValue {
	stateLen := len(state)
	newState := initEmptyState(stateLen)

	for i, row := range state {
		for j := range row {
			b := state[i][j]
			for k := 0; k < stateLen; k++ {
				b = b ^ f(state[k][j])
			}
			newState[j][i] = b
		}
	}

	return newState
}

func xor(a, b StateValue) StateValue {
	return a ^ b
}

func xorTable(state, other [][]StateValue) [][]StateValue {
	newState := initEmptyState(len(state))

	for i, row := range state {
		for j, val := range row {
			tmp := xor(val, other[i][j])
			newState[i][j] = tmp
		}
	}

	return newState
}

func generateRConstants(size int) [][][]StateValue {
	rConstants := make([][][]StateValue, size)

	var r StateValue
	var prevR StateValue

	for i := range rConstants {
		if i == 0 {
			r = StateValue(1)
		} else if prevR < StateValue(0x80) {
			r = prevR << 1
		} else if prevR >= StateValue(0x80) {
			r = (prevR << 1) ^ 0x1B
		}

		rConstants[i] = [][]StateValue{
			{0x00, 0x00},
			{r, 0x00},
		}

		prevR = r
	}

	return rConstants
}

func keyExpansion(key, size, rounds int, rConstants [][][]StateValue) [][][]StateValue {
	// Generate starting startingKeys from key input
	startingKeys := initEmptyState(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			r := (i + j + size) % key
			if r < 0 {
				r += size
			}
			startingKeys[i][j] = StateValue(r)
		}
	}

	// Concatenate keys into words
	expandedKeys := make([][][]StateValue, rounds)
	expandedKeys[0] = startingKeys
	for i := 0; i < rounds-1; i++ {
		expandedKeys[i+1] = xorTable(
			expandedKeys[i],
			xorTable(
				subBytes(rotWord(expandedKeys[i])),
				rConstants[i],
			),
		)
	}

	return expandedKeys
}

func rotWord(state [][]StateValue) [][]StateValue {
	stateLen := len(state)
	newState := initEmptyState(stateLen)

	for i, row := range state {
		for j := range row {
			if j == stateLen-1 {
				newState[j][i] = state[0][(i+1)%stateLen]
				continue
			}
			newState[j][i] = state[j+1][i]
		}
	}

	return newState
}

func addRoundKey(state [][]StateValue, keys [][]StateValue) [][]StateValue {
	return xorTable(state, keys)
}

func encrypt(state [][]StateValue, keys [][][]StateValue) [][]StateValue {
	//Rounds is derived from the number of keys
	rounds := len(keys) - 1

	// Initial round
	state = addRoundKey(state, keys[0])

	// Regular rounds
	for i := 1; i < rounds; i++ {
		state = subBytes(state)
		state = shiftRows(state)
		state = mixColumns(state)
		state = addRoundKey(state, keys[i])
	}

	// Final round
	state = subBytes(state)
	state = shiftRows(state)
	state = addRoundKey(state, keys[rounds])

	return state
}

func main() {
	rounds := 3
	size := 2
	key := 13

	fmt.Println("R_CONSTANTS")
	R_CONSTANTS := generateRConstants(rounds + 1)
	for _, r := range R_CONSTANTS {
		fmt.Println(tableToString(r))
	}

	fmt.Println("KEY_EXPANSION")
	keys := keyExpansion(key, size, rounds+1, R_CONSTANTS)
	for _, key := range keys {
		fmt.Println(tableToString(key))
	}

	fmt.Println("START STATE")
	state := [][]StateValue{
		{0x00, 0x02},
		{0x01, 0x03},
	}
	fmt.Println(tableToString(state))

	fmt.Println("Encrypted")
	fmt.Println(tableToString(encrypt(state, keys)))
}
