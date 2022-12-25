package main

import (
	"fmt"

	"github.com/thunderferret/TuckerGo/banking"
)

func main() {
	account := banking.Account{Owner: "nicloas", Balance: 20000}
	fmt.Println(account)
}
