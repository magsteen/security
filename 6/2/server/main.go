package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/pbkdf2"
)

const SALT_STR = "Salt"
const STORED_HASH = "ab29d7b5c589e18b52261ecba1d3a7e7cbf212c6"
const ITERATIONS = 2048
const SIZE = sha1.Size

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", login)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Printf("Error msg: %q", err)
	}
}

type Login struct {
	Hash string
}

func (l *Login) Verify() bool {
	serverHash := pbkdf2.Key([]byte(l.Hash), []byte(SALT_STR), ITERATIONS, SIZE, sha1.New)
	return hmac.Equal([]byte(serverHash), []byte(STORED_HASH))
}

func login(w http.ResponseWriter, r *http.Request) {
	var lgn Login
	err := json.NewDecoder(r.Body).Decode(&lgn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if lgn.Verify() {
		w.Header().Set("data", "Accepted")
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.Header().Set("data", "Failed")
		w.WriteHeader(http.StatusForbidden)
	}
}
