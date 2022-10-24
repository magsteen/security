package main

import "fmt"

var BIG_ALPH = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")
var LIL_ALPH = []rune("abcdefghijklmnopqrstuvwxyzæøå")

func main() {
	encrypt_a := 11
	encrypt_b := 5

	decrypt_a := 37
	decrypt_b := 11

	// a)
	kryptedAlphabet := enkrypt(LIL_ALPH, encrypt_a, encrypt_b)
	fmt.Printf("Krypt sequence: %s\n", kryptedAlphabet)

	// b)
	// funksjonen er surjektiv, som betyr at det eksisterer en input for alle mulige
	// output innenfor mengden brukt i funksjonen.

	// c)
	dekryptedAlphabet := dekrypt([]rune(kryptedAlphabet), decrypt_a, decrypt_b)
	fmt.Printf("Inverse sequence: %s\n", dekryptedAlphabet)

	m := "alice"
	fmt.Printf("'%s' enkrypted: %s\n", m, enkrypt([]rune(m), encrypt_a, encrypt_b))

	c := "SIØPBE"
	fmt.Printf("'%s' dekrypted: %s\n", c, dekrypt([]rune(c), decrypt_a, decrypt_b))
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

func enkrypt(mes []rune, a, b int) string {
	n := len(LIL_ALPH)

	var res string
	for _, m := range mes {
		i := index(LIL_ALPH, m)
		res += string(f(a, i, b, n))
	}

	return res
}

func dekrypt(mes []rune, a, b int) string {
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
