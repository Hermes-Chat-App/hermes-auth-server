package main

import (
	"fmt"

	route "github.com/Hermes-chat-App/hermes-auth-server/internal/router"
)

func main() {
	fmt.Println("Hello, world")
	route.InitRouter()
}
