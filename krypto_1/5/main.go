package main

import (
	"fmt"
	"strings"
)

var BIG_ALPH = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")
var LIL_ALPH = []rune("abcdefghijklmnopqrstuvwxyzæøå")
var ALPH_LEN = len(BIG_ALPH)

func main() {
	// Default test
	mes := "fisk"
	key := "lat"
	res := encrypt(mes, key)
	fmt.Printf("Start value: %s, Encrypted: %s, ", mes, res)

	res = decrypt(res, key)
	fmt.Printf("Decrypted: %s\n", res)

	// a)
	mesA := "Snart helg"
	keyA := "torsk"
	resA := encrypt(mesA, keyA)
	fmt.Printf("Start value: %s, Encrypted: %s\n", mesA, resA)

	// b)
	mesB := "QZQOBVCAFFKSDC"
	keyB := "brus"
	resB := decrypt(mesB, keyB)
	fmt.Printf("Start value: %s, Decrypted: %s\n", mesB, resB)

	// c)
	// a 15 character key gives 15^29 permutations (if only using the small alphabet).
	// 15^29 equals roughly 1.278340e+034
	// Will not consider this safe at all
}

func encrypt(mes, key string) string {
	var res string
	mes = strings.ToLower(mes)
	keyLen := len(key)
	keyIdx := -1
	for _, m := range mes {
		if m == ' ' {
			continue
		}
		keyIdx++
		keyIdx = keyIdx % keyLen
		mVal := index(LIL_ALPH, m)
		keyVal := index(LIL_ALPH, rune(key[keyIdx]))
		res += string(BIG_ALPH[(mVal+keyVal)%ALPH_LEN])
	}
	return res
}

func decrypt(mes, key string) string {
	var res string
	mes = strings.ToUpper(mes)
	keyLen := len(key)
	keyIdx := -1
	for _, m := range mes {
		if m == ' ' {
			continue
		}
		keyIdx++
		keyIdx = keyIdx % keyLen
		mVal := index(BIG_ALPH, m)
		keyVal := index(LIL_ALPH, rune(key[keyIdx]))
		r := (mVal - keyVal) % ALPH_LEN
		if r < 0 {
			r += ALPH_LEN
		}
		res += string(LIL_ALPH[r])
	}
	return res
}

func index(slice []rune, item rune) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	panic(fmt.Sprintf("item: %s not found in slice", string(item)))
}
