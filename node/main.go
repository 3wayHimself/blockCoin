package main

import (
	"bytes"
	"crypto/sha256"
	//"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
	//"net"
	//"flag"
)

/*
Dear Rob Pike,
	Go is a decent language, but for the sake of us 'Gophers', please use a better form of error handeling

Please don't kill me,
	Ender/The_Sushi/Josh
*/
func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

const targetBits = 21 //Max of 256, the lower the harder.  21 is both a prime number, and a triangular number.

type Block struct {
	Timestamp uint32
	Data      []byte
	PrevHash  []byte
	BlockHash []byte
	Nonce     uint32
}

type Blockchain []*Block

/* An unneeded function that creates a new blockchain with a Genesis block.
   I have no idea why I did this */
func NewBlockchain() Blockchain {
	return Blockchain{ GenesisBlock() }
}

//Creates a new Genesis block
func GenesisBlock() *Block {
	return NewBlock( "GENISIS", []byte{} )
}

//Adds a block to the blockchain
func (bc Blockchain) AddBlock(data string) {
	previousBlock := bc[ len(bc)-1 ]
	newBlock := NewBlock(data, previousBlock.BlockHash)
	bc = append(bc, newBlock)
}

//Creates a new block
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{ uint32(time.Now().Unix()), []byte(data), prevHash, []byte{}, 0 }
	hash, nonce := MineBlock(block)
	block.BlockHash = hash
	block.Nonce = uint32(nonce)
	return block
}

//Converts a uint32 to a byte array
func ToBytes(num uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, num)
	return bs
}

/*
	MineBlock(block *Block) ([]byte, uint32) -- Mines a block.
	----------------------------------------------------------
	Alright, let's go through the algorithm here:
		First, we create a new bigInt with a value of 1, then we left-shift that by 256 minus the target bits.
		The reason for this, is that we can't have the target be larger than 256 bits, or it will be larger than the SHA256 hash.
		We then put targetBits in to a variable called targetBitsArray.  After this, we join all the block data into a 2d byte array,
		then we begin the mining process.
		In mining, we start the nonce at 0, then move up.  We hash the data, then check if the hash is less than the target.
		If not, we increment the nonce, and repeat the same process all over again. Once the hash is less than the target, we return the hash and nonce.
*/
func MineBlock(block *Block) ([]byte, uint32) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) //The SHA256 hash should take 256 bits of memory, we left-shift by the size of target
	targetBitsArray := ToBytes(targetBits)

	data := bytes.Join(
		[][]byte{
			ToBytes(block.Timestamp),
			block.Data,
			block.PrevHash,
			targetBitsArray,
			ToBytes(uint32(block.Nonce)),
		},
		[]byte{},
	)

	var hashInt big.Int
	var hash [32]byte
	var nonce uint32
	nonce = 0
	hash = sha256.Sum256(data)
	hashInt.SetBytes(hash[:]) //we need it as a slice, so we have [:]
	fmt.Printf("Mining block containing '%s'\n", block.Data)
	for hashInt.Cmp(target) != -1 { //we should check for a nonce overflow, although rare, it is possible
		nonce++
		block.Nonce = nonce
		data := bytes.Join(
			[][]byte{
				ToBytes(block.Timestamp),
				block.Data,
				block.PrevHash,
				targetBitsArray,
				ToBytes(uint32(block.Nonce)),
			},
			[]byte{},
		)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
	}

	fmt.Println("Mined!")
	return hash[:], nonce
}

/*
func Serv(client net.Conn) {
}
*/

func main() {
	/*
	fmt.Println("Starting server...")
	listener, err := net.Listen("tcp", ":1912")
	Handle(err)
	for {
		client, err := listener.Accept()
		Handle(err)
		go Serv(client)
	}
	*/
	
	chain := NewBlockchain()
	chain.AddBlock("Here is some test data.")
	chain.AddBlock("Later, we'll use transaction data instead")
}
