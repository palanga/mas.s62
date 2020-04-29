package main

import (
    "time"
    "math/rand"
    "bytes"
)

// This file is for the mining code.
// Note that "targetBits" for this assignment, at least initially, is 33.
// This could change during the assignment duration!  I will post if it does.

// Mine mines a block by varying the nonce until the hash has targetBits 0s in
// the beginning.  Could take forever if targetBits is too high.
// Modifies a block in place by using a pointer receiver.
func (self EventualBlock) Mine(targetBits uint8) Block {
	// your mining code here
	// also feel free to get rid of this method entirely if you want to
	// organize things a different way; this is just a suggestion

	 nonceLenght := 100 - 64 - 1 - len(self.Name) - 1

	for !CheckWork(self, targetBits) {
		self.Nonce = GenerateRandomNonce(nonceLenght)
	}

	var block Block

	block.PrevHash = self.PrevHash
	block.Name = self.Name
	block.Nonce = self.Nonce

	return block
}

// CheckWork checks if there's enough work
func CheckWork(eventualBlock EventualBlock, targetBits uint8) bool {
	// your checkwork code here
	// feel free to inline this or do something else.  I just did it this way
	// so I'm giving empty functions here.

	return eventualBlock.Hash().StartsWithZeros(targetBits)
}

// TODO hardcoded 0000
func (hash Hash) StartsWithZeros(targetBits uint8) bool {
	return bytes.HasPrefix([]byte(hash.ToString()), []byte("0000"))
}


var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func GenerateRandomNonce(lenght int) string {
    rand.Seed(time.Now().UnixNano())

    return randSeq(lenght)
}
