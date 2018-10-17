package main

import (
	"fmt"

	"github.com/swp/access/configs"
)

func main() {
	var pass string = "kiki123"
	var passHash string = fmt.Sprintf("%x", configs.Encrypt(pass))

	fmt.Println(fmt.Sprintf("ORIGINAL %s", pass))
	fmt.Println(fmt.Sprintf("HASHED  %s", passHash))
}
