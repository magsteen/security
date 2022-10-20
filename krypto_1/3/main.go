package main

import "fmt"

var BIG_ALPH = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")
var LIL_ALPH = []rune("abcdefghijklmnopqrstuvwxyzæøå")

func main() {
	kryptedAlphabet := enkrypt(LIL_ALPH)
	fmt.Printf("Krypt sequence: %s\n", kryptedAlphabet)

	dekryptedAlphabet := dekrypt([]rune(kryptedAlphabet))
	fmt.Printf("Inverse sequence: %s\n", dekryptedAlphabet)

	m := "alice"
	fmt.Printf("'%s' enkrypted: %s\n", m, enkrypt([]rune(m)))

	c := "SIØPBE"
	fmt.Printf("'%s' dekrypted: %s\n", c, dekrypt([]rune(c)))
}

func f(a, x, b, n int) rune {
	r := (a*x - b) % n
	if r < 0 {
		r += n
	}
	return BIG_ALPH[r]
}

func fInvers(a, y, b, n int) rune {
	r := (a*y + b) % n
	if r < 0 {
		r += n
	}
	return LIL_ALPH[r]
}

func enkrypt(mes []rune) string {
	a := 11
	b := 5
	n := len(LIL_ALPH)

	var res string
	for _, m := range mes {
		i := index(LIL_ALPH, m)
		res += string(f(a, i, b, n))
	}

	return res
}

func dekrypt(mes []rune) string {
	a := 37
	b := 11
	n := len(BIG_ALPH)

	var res string
	for _, m := range mes {
		i := index(BIG_ALPH, m)
		res += string(fInvers(a, i, b, n))
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
