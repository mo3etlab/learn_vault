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
		// The problem here is that the type of middle.GetMiddleware("middleware1") is middle.Middleware,
		// but the Chain function requires a parameter of type Middleware.
		// We need to ensure type consistency. If the two types have the same definition,
		// we can use them directly. Otherwise, we need to convert them.
		// Assume the types are the same, and we can use them directly as follows:
		// 这里的问题是，middle.GetMiddleware("middleware1") 的返回类型是 middle.Middleware，
		// 但 Chain 函数需要的参数类型是 Middleware。
		// 我们需要确保类型一致。如果这两种类型定义相同，
		// 我们可以直接使用它们。否则，我们需要进行类型转换。
		// 假设类型相同，我们可以直接按如下方式使用：

		// Because have different type name, we need to convert it
		Middleware(middle.GetMiddleware("middleware1")),
		Middleware(middle.GetMiddleware("middleware2")),
		Middleware(middle.GetMiddleware("middleware3")),
	)

	// execute handle progress
	result := handler("hello world")
	fmt.Println("\n Final result:", result)
}
