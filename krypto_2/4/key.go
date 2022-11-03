package main

import (
	"math/rand"
	"time"
)

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

		rConstants[i] = initEmptyState(size)
		rConstants[i][0][0] = r

		prevR = r
	}

	return rConstants
}

func createKeyFromValues(stateSize int, keyValues ...StateValue) [][]StateValue {
	if stateSize*stateSize != len(keyValues) {
		panic("stateSize mismatch the amount of key words")
	}
	key := initEmptyState(stateSize)

	for i := 0; i < stateSize; i++ {
		colValues := keyValues[stateSize*i : stateSize*(i+1)]
		for j := 0; j < stateSize; j++ {
			key[j][i] = colValues[j]
		}
	}

	return key
}

func createRandomKey(stateSize int) [][]StateValue {
	rand.Seed(time.Now().UnixNano())
	key := initEmptyState(stateSize)

	for i := 0; i < stateSize; i++ {
		for j := 0; j < stateSize; j++ {
			key[j][i] = StateValue(rand.Intn(255 + 1))
		}
	}
	return key
}

func rotWord(word [][]StateValue) [][]StateValue {
	wordSize := len(word)
	newWord := make([][]StateValue, wordSize)

	for i, row := range word {
		newWord[(i+1)%wordSize] = row
	}

	return newWord
}

func keyExpansion(mainKey [][]StateValue, rounds int) [][][]StateValue {
	colLen := len(mainKey)
	N := len(mainKey[0])
	expandedKeys := make([][][]StateValue, rounds)

	WiN := make([][]StateValue, colLen)
	prevW := make([][]StateValue, colLen)
	word := make([][]StateValue, colLen)

	counter := 0
	for i := 0; i < rounds-1; i++ {
		roundKey := make([][]StateValue, N)

		for j := 0; j < N; j++ {
			if counter < N {
				word = getTableColumn(mainKey, j)
			} else {
				WiN = getTableColumn(expandedKeys[i-1], j)

				if counter >= N && counter%N == 0 {
					word = xorTable(
						xorTable(
							WiN,
							subBytes(rotWord(prevW)),
						),
						R_CONSTANTS[counter/N],
					)
				} else if counter >= N && N > 6 && counter%N == 4 {
					word = xorTable(WiN, subBytes(prevW))
				} else {
					word = xorTable(WiN, prevW)
				}
			}

			roundKey = appendTableColumn(roundKey, word)
			prevW = word
			counter++
		}
		expandedKeys[i] = roundKey
	}

	return expandedKeys
}
