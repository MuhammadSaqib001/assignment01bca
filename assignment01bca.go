package assignment01bca
import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type block struct {
	nounce        int
	data         string
	previousHash string
	currentHash  string
}

type Blockchain struct {
	block_lists []block
}

func CalculateHash(stringToHash string) (output string) {
	output = fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
	return
}

func NewBlock(transaction string, nounce int, previousHash string) (bl *block) {
	new_block := block{nounce: nounce, data: transaction, previousHash: previousHash}
	a = &new_block
	return a
}

func Change_block(my_new *block) {
	previous_nounce := my_new.nounce
	my_new.nounce = rand.Intn(1000)
	if my_new.nounce == previous_nounce {
		my_new.nounce = rand.Intn(1000)
	}
	hash := CalculateHash(strconv.Itoa(b.nounce) + my_new.data)
	my_new.currentHash = hash
}

func Display_blocks(chain []block) {
	i := 0
	for i < len(chain) {
		fmt.Printf("\n\nBlock_id: %d\n Nounce: %d\nData: %v\nPrevious Hash: %v\nCurrent Hash: %v", i, chain[i].nounce, chain[i].data, chain[i].previousHash, chain[i].currentHash)
		fmt.Println()
		i++
	}
}

func VerifyChange(chain []block) {
	i := 0
	for i < len(chain)-1 {
		if chain[i].currentHash != chain[i+1].previousHash {
			fmt.Printf("\n\n Blocks were modified \nBlock id: %d\n\n", i)
		}
		i++
	}
}
