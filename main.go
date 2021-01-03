package main

import (
	"fmt"
	oauth "gitapp/oauth"
	user "gitapp/user"
)

func main() {
	fmt.Println("Starting main function")
	go user.Handler()
	oauth.Handler()
}
