package main

import "fmt"

var R_CONSTANTS = []State{
	{{0x01}, {0x00}, {0x00}, {0x00}},
	{{0x02}, {0x00}, {0x00}, {0x00}},
	{{0x04}, {0x00}, {0x00}, {0x00}},
	{{0x08}, {0x00}, {0x00}, {0x00}},
	{{0x10}, {0x00}, {0x00}, {0x00}},
	{{0x20}, {0x00}, {0x00}, {0x00}},
	{{0x40}, {0x00}, {0x00}, {0x00}},
	{{0x80}, {0x00}, {0x00}, {0x00}},
	{{0x1b}, {0x00}, {0x00}, {0x00}},
	{{0x36}, {0x00}, {0x00}, {0x00}},
	{{0x6c}, {0x00}, {0x00}, {0x00}},
	{{0xd8}, {0x00}, {0x00}, {0x00}},
	{{0xab}, {0x00}, {0x00}, {0x00}},
	{{0x4d}, {0x00}, {0x00}, {0x00}},
	{{0x9a}, {0x00}, {0x00}, {0x00}},
	{{0x2f}, {0x00}, {0x00}, {0x00}},
}

func main() {
	size := 4
	rounds := 14
	msg := "Dette er en text" + "En annen text"

	states := textToStateBlocks(msg, size)
	key := createRandomKey(size)

	fmt.Println("KEY_EXPANSION")
	keys := keyExpansion(key, rounds+1)
	for _, key := range keys {
		fmt.Println(key.tableToString())
	}

	fmt.Println("START STATE")
	fmt.Printf("text: %s\n", msg)
	fmt.Println("state block(s):")
	for _, state := range states {
		fmt.Println(state.tableToString())
	}

	fmt.Println("END STATE (ENCRYPTED)")
	encryptedStates := encrypt(msg, size, keys)
	fmt.Printf("text: %s\n", stateBlocksToText(encryptedStates))
	fmt.Println("state block(s):")
	for _, encryptedState := range encryptedStates {
		fmt.Println(encryptedState.tableToString())
	}
}
