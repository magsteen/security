package main

import (
	"fmt"
)

var ALPH = []rune(" abcdefghijklmnopqrstuvwxyzæøå,.")

const IV = 13

var RUNE_MAP, BYTE_MAP = createAlphabetMap()
var ALPH_LEN = len(ALPH)

// func intToBitString(i int) byte {
// 	return byte(fmt.Sprintf("%05b", i))
// }

func createAlphabetMap() (map[rune]int, map[int]rune) {
	runeMap := make(map[rune]int)
	byteMap := make(map[int]rune)

	for i, symbol := range ALPH {
		runeMap[symbol] = i
		byteMap[i] = symbol
	}
	return runeMap, byteMap
}

func main() {
	// a)
	res := encrypt("dette er en test", 5)
	fmt.Println(res)
}

func xor(a, b int) int {
	return int(byte(a) ^ byte(b))
}

// Cipher table
func cesarEncrypt(x int) int {
	r := (x + 3) % ALPH_LEN
	if r < 0 {
		r += ALPH_LEN
	}
	return r
}

// Dechiper table
func cesarDecrypt(x int) int {
	r := (x - 3) % ALPH_LEN
	if r < 0 {
		r += ALPH_LEN
	}
	return r
}

func padString(s string, l int) string {
	for i := len(s); i < l; i++ {
		s += "."
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

func concatBlocks(blocks []string) string {
	var result string
	for _, block := range blocks {
		result += block
	}

	return result
}

func encrypt(msg string, blockLen int) string {
	var newBlocks []string
	p := make([]int, blockLen)
	newP := p

	// init p for xor
	for i := 0; i < blockLen; i++ {
		p[i] = IV
	}

	blocks := splitToBlocks(msg, blockLen)
	for _, block := range blocks {
		var symAsInt int
		var newInt int
		var newBlock string
		for i, symbol := range block {
			symAsInt = RUNE_MAP[symbol]
			newInt = cesarEncrypt(xor(p[i], symAsInt))
			newBlock += string(BYTE_MAP[newInt])
			newP[i] = newInt
		}

		p = newP
		newBlocks = append(newBlocks, newBlock)
	}

	return concatBlocks(newBlocks)
}
