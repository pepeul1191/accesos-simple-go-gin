package main

import (
	"fmt"
	"log"

	"github.com/swp/access/configs"
)

func main() {
	configs.SetConstants()
	//var passHash string = fmt.Sprintf("%x", configs.Encrypt([]byte(pass)))
	pass := "kiki123"
	fmt.Printf("%s\n", pass)
	ciphertext, err := configs.Encrypt(pass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0x\n", ciphertext)
	result, err := configs.Decrypt(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}
