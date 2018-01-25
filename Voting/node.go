package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

const targetBits = 32

type Block struct {
	Timestamp uint32
	Data      []byte
	PrevHash  []byte
	BlockHash []byte
	Nonce     uint32
}

type Blockchain []*Block

//Creates a new blockchain with a Genesis block
func NewBlockchain() Blockchain {
	return Blockchain{ GenesisBlock() }
}

//Creates a new Genesis block
func GenesisBlock() *Block {
	return NewBlock( "Genesis Block", []byte{} )
}

func (bc Blockchain) AddBlock(data string) {
	previousBlock := bc[ len(bc)-1 ] //This is to retrieve the previous block hash so you don't have to later on
	newBlock := NewBlock(data, previousBlock.BlockHash)
	bc = append(bc, newBlock)
}

//Made it so the data param is a string
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{ uint32(time.Now().Unix()), []byte(data), prevHash, []byte{}, 0 }
	hash, nonce := MineBlock(block)
	block.BlockHash = hash[:]
	block.Nonce = uint32(nonce)
	return block
}

func ToBytes(num uint32, bits int8) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, num)
	return bs
}

func MineBlock(block *Block) ([]byte, uint32) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, targetBits)

	data := bytes.Join(
		[][]byte{
			ToBytes(block.Timestamp, 32),
			block.Data,
			block.PrevHash,
			bs,
			ToBytes(uint32(block.Nonce), 32),
		},
		[]byte{},
	)

	var hashInt big.Int
	var hash [32]byte
	var nonce uint32
	nonce = 0
	hash = sha256.Sum256(data)
	fmt.Printf("Mining block containing '%s'\n", block.Data) //Have to use double quotes for \n to work
	for hashInt.Cmp(target) != -1 {
		nonce++
		block.Nonce = nonce
		data := bytes.Join(
			[][]byte{
				ToBytes(block.Timestamp, 32),
				block.Data,
				block.PrevHash,
				bs,
				ToBytes(uint32(block.Nonce), 32),
			},
			[]byte{},
		)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
	}

	fmt.Println("Mined!")
	return hash[:], nonce
}

func main() {
	chain := NewBlockchain()
	chain.AddBlock("This is a test")
}
