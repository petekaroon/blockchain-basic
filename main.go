package main

import "fmt"

func main() {
	blockchainOne := createBlockchain(2)

	blockchainOne.addBlock("Alice", "Bob", 5)
	blockchainOne.addBlock("John", "Bob", 2)
	blockchainOne.addBlock("Alice", "John", 3)

	fmt.Println(blockchainOne.isValid())
}
