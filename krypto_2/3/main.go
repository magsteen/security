package main

import (
	"fmt"
	"strings"
)

const ALPH_STRING = "abcdefghijklmnopqrstuvwxyzæøå"

var ALPH_RUNES = []rune(ALPH_STRING)
var ALPH_LEN = len(ALPH_RUNES)

var RUNE_MAP, INT_MAP = createAlphabetMap()

func createAlphabetMap() (map[rune]int, map[int]rune) {
	runeMap := make(map[rune]int)
	byteMap := make(map[int]rune)

	for i, symbol := range ALPH_RUNES {
		runeMap[symbol] = i
		byteMap[i] = symbol
	}
	return runeMap, byteMap
}

func main() {
	k1 := 17
	plain1 := "goddag"
	crypted1 := encrypt(plain1, k1)
	fmt.Printf("Encrypt(%s) -> %v\n", plain1, crypted1)

	k2 := 5
	crypted2 := []int{23, 8, 23, 12, 21, 02, 04, 3, 17, 13, 19}
	plain2 := decrypt(crypted2, k2)
	fmt.Printf("Decrypt(%v) -> %s\n", crypted2, plain2)
}

func createKeyStream(s string, k int) []int {
	res := make([]int, len(s))
	runes := []rune(s)

	res[0] = k
	for i := 0; i < len(s)-1; i++ {
		res[i+1] = RUNE_MAP[runes[i]]
	}
	return res
}

func skift(P, Z int) int {
	C := (P + Z) % ALPH_LEN
	if C < 0 {
		C += ALPH_LEN
	}

	return C
}

func unskift(C, Z int) int {
	P := (C - Z) % ALPH_LEN
	if P < 0 {
		P += ALPH_LEN
	}

	return P
}

func encrypt(msg string, k int) []int {
	res := make([]int, len(msg))
	keyStream := createKeyStream(msg, k)
	keyStreamLen := len(keyStream)
	var P, Z int

	for i, PRune := range msg {
		P = RUNE_MAP[PRune]
		Z = keyStream[i%keyStreamLen]
		res[i] = skift(P, Z)
	}

	return res
}

func decrypt(cipherStream []int, k int) string {
	var sb strings.Builder
	Z := k

	for _, C := range cipherStream {
		P := unskift(C, Z)

		if _, err := sb.WriteRune(INT_MAP[P]); err != nil {
			panic(err)
		}

		Z = P
	}

	return sb.String()
}
