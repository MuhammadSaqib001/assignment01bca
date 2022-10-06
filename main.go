package main

import (
	"fmt"
	"math/rand"
	"strconv"

	git "github.com/MuhammadSaqib001/assignment01bca"
)

func main() {
	my_blockchain := []git.block{}
	key := 0
	fmt.Print("Do you want to add new block \n 1. Yes \n2. No\n")
	fmt.Scan(&key)
	previous_hash := ""
	for key == 1 {
		coming_block := git.NewBlock("SAQIB to RASIB", rand.Intn(100), previous_hash)
		previous_hash = git.CalculateHash(strconv.Itoa(coming_block.nonce) + coming_block.data)
		coming_block.currentHash = previous_hash
		my_blockchain = append(my_blockchain, *coming_block)

		fmt.Print("Want to add new block \n1. Yes 2. No\n")
		fmt.Scan(&key)
	}

	git.Display_blocks(my_blockchain)
	git.Change_block(&my_blockchain[2])
	git.VerifyChange(my_blockchain)
}
