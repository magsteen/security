package main

import "fmt"

var BIG_ALPH = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")
var LIL_ALPH = []rune("abcdefghijklmnopqrstuvwxyzæøå")
var n = len(BIG_ALPH)

func main() {
	mes := "fisk"
	key := "lat"
	res := encrypt(mes, key)
	fmt.Printf("Encrypted: %s\n", res)

	res = decrypt(res, key)
	fmt.Printf("Decrypted: %s\n", res)
}

func encrypt(mes, key string) string {
	var res string
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
		res += string(BIG_ALPH[(mVal+keyVal)%n])
	}
	return res
}

func decrypt(crypted, key string) string {
	var res string
	keyLen := len(key)
	keyIdx := -1
	for _, m := range crypted {
		if m == ' ' {
			continue
		}
		keyIdx++
		keyIdx = keyIdx % keyLen
		mVal := index(LIL_ALPH, m)
		keyVal := index(LIL_ALPH, rune(key[keyIdx]))
		res += string(BIG_ALPH[(mVal+keyVal)%n])
	}
	return res
}

func index(slice []rune, item rune) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	panic("item not found in slice")
}
