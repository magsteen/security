package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/pbkdf2"
)

const SALT_STR = "Salt"

var STORED_HASH = []byte{0x00, 0x07, 0xbf, 0xb0, 0xa9, 0xb6, 0x75, 0xfa, 0xfd,
	0x60, 0x84, 0x80, 0xa3, 0x19, 0xe9, 0xa4, 0xa4, 0x1e, 0xa3, 0x1f}

const ITERATIONS = 2048
const SIZE = sha1.Size
const PORT = ":3000"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login)

	fmt.Printf("Start listening on %s\n", PORT)
	err := http.ListenAndServe(PORT, cors.AllowAll().Handler(r))
	if err != nil {
		fmt.Printf("Error msg: %q", err)
	}
}

type LoginReq struct {
	Hash string `json:"hash"`
}

type LoginRes struct {
	Ok bool `json:"ok"`
}

func (l *LoginReq) Verify() bool {
	serverHash := pbkdf2.Key([]byte(l.Hash), []byte(SALT_STR), ITERATIONS, SIZE, sha1.New)
	fmt.Printf("Hash of recieved: %x\nHash of stored:   %x\n", serverHash, STORED_HASH)
	return hmac.Equal([]byte(serverHash), []byte(STORED_HASH))
}

func login(w http.ResponseWriter, r *http.Request) {
	var lgn LoginReq
	err := json.NewDecoder(r.Body).Decode(&lgn)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	if lgn.Verify() {
		p := LoginRes{true}
		json.NewEncoder(w).Encode(p)
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.Header().Set("data", "Failed")
		w.WriteHeader(http.StatusForbidden)
	}
}
