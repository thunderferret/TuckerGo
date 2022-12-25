package main

import (
	"./banking"

	"fmt"
)

func main() {
	account := banking.Account{Owner: "nicloas", Balance: 20000}
	fmt.Println(account)
}
