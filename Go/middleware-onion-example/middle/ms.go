package middle

import (
	"fmt"
	"time"
)

// core handler function type
type Handler func(string) string

func middleware1(next Handler) Handler {
	return func(input string) string {
		// fmt.Println("middleware1 处理中...")
		fmt.Println("[中间件1] 预处理 - 添加前缀")
		modified := "prefix_ " + input

		// return next(input)

		result := next(modified)

		// fmt.Println("middleware1 处理结束")
		fmt.Println("[中间件1] 后处理")

		return result
	}
}

func middleware2(next Handler) Handler {
	return func(input string) string {
		fmt.Println("[中间件2] 预处理 - 开始计时")
		start := time.Now()

		result := next(input)

		fmt.Printf("[中间件2] 后处理 - 耗时: %v\n", time.Since(start))
		return result

	}
}
