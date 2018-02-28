package main

import (
	"fmt"
)

var (
	PROGRAM     string
	VERSION     string
	BUILD       string
	COMMIT_SHA1 string
)

func init() {
	fmt.Printf("-----------------------------------------------------\nProgram: %s\nVersion: %s\nBuild: %s\nCommit_sha1: %s\n-----------------------------------------------------\n", PROGRAM, VERSION, BUILD, COMMIT_SHA1)
}

// func main() {
// 	bc := NewBlockchain()

// 	bc.AddBlock("Send 1 BTC to Ivan")
// 	bc.AddBlock("Send 2 more BTC to Ivan")

// 	for _, block := range bc.blocks {
// 		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
// 		fmt.Printf("Data: %s\n", block.Data)
// 		fmt.Printf("Hash: %x\n", block.Hash)
// 		fmt.Printf("Nonce: %v\n", block.Nonce)
// 		pow := NewProofOfWork(block)
// 		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
// 		fmt.Println()
// 	}

// }

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
