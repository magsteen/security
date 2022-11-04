package main

import (
	"math/rand"
	"time"
)

func createRandomKey(stateSize int) State {
	rand.Seed(time.Now().UnixNano())
	key := NewZeroedState(stateSize, stateSize)

	for i := 0; i < stateSize; i++ {
		for j := 0; j < stateSize; j++ {
			key[j][i] = byte(rand.Intn(255 + 1))
		}
	}
	return key
}

func rotWord(word State) State {
	wordSize := len(word)
	newWord := make(State, wordSize)

	for i, row := range word {
		newWord[(i+wordSize-1)%wordSize] = row
	}

	return newWord
}

func keyExpansion(mainKey State, rounds int) []State {
	N := len(mainKey[0])
	expandedKeys := make([]State, rounds)

	var WiN State
	var prevW State
	var word State

	counter := 0
	for i := 0; i < rounds-1; i++ {
		roundKey := make(State, N)

		for j := 0; j < N; j++ {
			if counter < N {
				word = mainKey.getTableColumn(j)
			} else {
				WiN = expandedKeys[i-1].getTableColumn(j)

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

			roundKey = roundKey.appendColumn(word)
			prevW = word
			counter++
		}
		expandedKeys[i] = roundKey
	}

	return expandedKeys
}
