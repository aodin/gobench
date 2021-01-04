package main

import (
	"fmt"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Int63() % int64(len(letters))]
    }
    return string(b)
}

func main() {
	fmt.Println(RandStringBytesRmndr(12))
}
