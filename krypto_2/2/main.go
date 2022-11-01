package main

import (
	"fmt"
	"strings"
)

const ALPH_STRING = " abcdefghijklmnopqrstuvwxyzæøå,."
const IV = 13

var ALPH_RUNES = []rune(ALPH_STRING)
var RUNE_MAP, INT_MAP = createAlphabetMap()
var ALPH_LEN = len(ALPH_RUNES)

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
	plain1 := "aaaaaa"
	crypted1 := encrypt(plain1)
	fmt.Printf("Encrypt(%s) -> %s\n", plain1, crypted1)

	plain2 := "dette er en test"
	crypted2 := encrypt(plain2)
	fmt.Printf("Encrypt(%s) -> %s\n", plain2, crypted2)

	crypted3 := "qvxæyy hkgdgk,,oqhdnc"
	plain3 := decrypt(crypted3)
	fmt.Printf("Decrypt(%s) -> %s\n", crypted3, plain3)
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

// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Cipher_block_chaining_(CBC)
func encrypt(msg string) string {
	var sb strings.Builder
	var newC int

	prevC := IV

	for _, P := range msg {
		newC = cesarEncrypt(xor(RUNE_MAP[P], prevC))

		_, err := sb.WriteRune(INT_MAP[newC])
		if err != nil {
			panic(err)
		}

		prevC = newC
	}

	return sb.String()
}

func decrypt(cipherText string) string {
	var sb strings.Builder
	var C int
	prevC := IV

	for _, CRune := range cipherText {
		C = RUNE_MAP[CRune]

		_, err := sb.WriteRune(INT_MAP[xor(cesarDecrypt(C), prevC)])
		if err != nil {
			panic(err)
		}

		prevC = C
	}

	return sb.String()
}
