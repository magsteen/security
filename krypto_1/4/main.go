package main

import "fmt"

var BIG_ALPH = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")
var LIL_ALPH = []rune("abcdefghijklmnopqrstuvwxyzæøå")
var ALPH_LEN = len(BIG_ALPH)

func main() {
	// Ved skift nr 12 ser vi en lesbar mld
	// 'hjerneneralene' -> 'hjernen er alene'

	c := "YÆVFB VBVFR ÅVBV"
	decrypt(c)
}

func decrypt(mes string) {
	var res string
	for i := 0; i < 29; i++ {
		res = ""
		for _, m := range mes {
			if m == ' ' {
				continue
			}
			j := index(BIG_ALPH, m)
			res += string(LIL_ALPH[(j+i)%ALPH_LEN])
		}
		fmt.Printf("Skift %d: %s\n", i, res)
	}
}

func index(slice []rune, item rune) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	panic("item not found in slice")
}
