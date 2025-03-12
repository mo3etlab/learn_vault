package main

import (
	"fmt"
	"strings"

	"github.com/mo3etlab/learn_vault/middleware-onion-example/middle"
)

// middleware type(decorator)
type Middleware func(middle.Handler) middle.Handler

// onion model middleware chain
func Chain(h middle.Handler, ms ...Middleware) middle.Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i](h)
	}
	return h
}

// core Business logic
func coreLogic(input string) string {
	fmt.Println("核心逻辑处理中...")
	return strings.ToUpper(input)
}

func main() {
	handler := Chain(
		coreLogic,
		// middleware1,
		// middleware2,
		// middleware3,
	)

	// execute handle progress
	result := handler("hello world")
	fmt.Println("\n Final result:", result)
}
