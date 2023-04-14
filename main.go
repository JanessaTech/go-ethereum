// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"fmt"
	gobinding "hi-supergirl/go-ethereum/go-binding"
)

func testGoBinding() {
	fmt.Println("start in main .....")
	//gobinding.DeployContract()
	gobinding.CallRetrieve()
	//gobinding.SendTransaction()
	//gobinding.MockTest() // need to fix bug: always complaining "Failed to deploy new storage contract: transaction type not supported"
}

func main() {
	testGoBinding()
}
