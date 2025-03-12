package middle

import (
	"fmt"
	"time"
)

// core handler function type
type Handler func(string) string

type Middleware func(Handler) Handler

func GetMiddleware(name string) Middleware {
	switch name {
	case "middleware1":
		return middleware1
	case "middleware2":
		return middleware2
	case "middleware3":
		return middleware3
	default:
		return func(next Handler) Handler {
			return func(input string) string {
				return next(input)
			}
		}
	}
}

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

func middleware3(next Handler) Handler {
	return func(input string) string {
		fmt.Println("[中间件3] 预处理 - 检查长度")
		if len(input) > 100 {
			return "错误：输入过长"
		}

		result := next(input)

		fmt.Println("[中间件3] 后处理 - 添加后缀")
		return result + "_suffix"
	}
}
