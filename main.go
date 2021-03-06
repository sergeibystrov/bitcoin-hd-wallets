package main

import (
	"fmt"

	"git.roosoft.com/bitcoin/hd-wallets/lib"

	seed "git.roosoft.com/bitcoin/hd-wallets/1-seed"
	masterkey "git.roosoft.com/bitcoin/hd-wallets/2-masterPrivateKey"
	childkeys "git.roosoft.com/bitcoin/hd-wallets/3-childKeys"
	xpub "git.roosoft.com/bitcoin/hd-wallets/4-xpub"
)

func doSeed() {
	fmt.Printf("\n\n1. Seed\n")
	fmt.Printf("-------\n")
	fmt.Printf("%-9s %s\n", "mnemonic", seed.GetMnemonic())
	fmt.Printf("%-9s %x\n", "seed", seed.GetSeed())
}

func doMasterKey() {
	fmt.Printf("\n\n2. Master key\n")
	fmt.Printf("-------------\n")
	fmt.Printf("%-12s %x\n", "private key", masterkey.GetPrivateKey())
	fmt.Printf("%-12s %x\n", "chain code", masterkey.GetChainCode())
	fmt.Printf("%-12s %x\n", "public key", masterkey.GetPublicKey())
}

func doChildKeys() {
	fmt.Printf("\n\n3. Child keys\n")
	fmt.Printf("-------------\n\n")
	fmt.Printf("---BIP 44 ---\n")

	// m/44h/0h/0h/0/*
	for i := uint32(0); i < 20; i++ {
		path, address, segwit32, segwitNested, _ := childkeys.GetChild(lib.PurposeBIP44, lib.CoinTypeBTC, 0, 0, i)
		fmt.Printf("%-16s %-34s %s %s\n", path, address, segwitNested, segwit32)
	}

	fmt.Printf("\n\n---BIP 84 ---\n")

	// m/84h/0h/0h/0/*
	for i := uint32(0); i < 20; i++ {
		path, address, segwit32, segwitNested, _ := childkeys.GetChild(lib.PurposeBIP84, lib.CoinTypeBTC, 0, 0, i)
		fmt.Printf("%-16s %-34s %-34s %-43s\n", path, address, segwitNested, segwit32)
	}
}

func doXpub() {
	fmt.Printf("\n\n4. xpub\n")
	fmt.Printf("-------\n\n")

	xpriv := xpub.GetXpriv(lib.PurposeBIP84, lib.CoinTypeBTC, 0, 0, 0)
	xpub := xpub.GetXpub(lib.PurposeBIP84, lib.CoinTypeBTC, 0, 0, 0)

	fmt.Printf("%-6s %s\n", "xpriv", xpriv)
	fmt.Printf("%-6s %s\n", "xpub", xpub)
}

func main() {
	doSeed()
	doMasterKey()
	doChildKeys()
	doXpub()
}
