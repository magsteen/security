package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	salt := []byte("Saltet til Ola")
	keyHash := []byte{0xab, 0x29, 0xd7, 0xb5, 0xc5, 0x89, 0xe1, 0x8b, 0x52, 0x26,
		0x1e, 0xcb, 0xa1, 0xd3, 0xa7, 0xe7, 0xcb, 0xf2, 0x12, 0xc6}
	iterations := 2048
	size := sha1.Size
	digest := sha1.New

	res := BruteForce([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"), keyHash, salt, 3, iterations, size, digest)
	fmt.Printf("Result: %c\n", res)
}

func BruteForce(alphabet, key, salt []byte, l, iterations, size int, digest func() hash.Hash) []byte {
	alphabetLen := len(alphabet)
	b := bytes.Repeat(alphabet[:1], l)
	prog := make([]int, l)

	for pos := l - 1; pos > 0; pos-- {
		for charInc := 0; charInc <= alphabetLen-1; charInc++ {
			if pos != l-1 {
				pos++
				charInc--
				continue
			}

			b[pos] = alphabet[charInc]
			hash := pbkdf2.Key(b, salt, iterations, size, digest)
			if hmac.Equal([]byte(hash), []byte(key)) {
				return b
			}

			prog[pos]++
		}
	}
	return nil
}
