package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := ":32880"
	log.Printf("Starting webhook server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// 记录请求基本信息
	log.Printf("Received %s request to %s", r.Method, r.URL.Path)

	// 读取请求头
	headers := make(map[string]string)
	for k, v := range r.Header {
		headers[k] = v[0]
	}

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// 格式化输出日志
	log.Println("=== Webhook Received ===")
	log.Printf("- Method: %s\n", r.Method)
	log.Printf("- Headers:\n")
	headerData, _ := json.MarshalIndent(headers, "", "  ")
	log.Println(string(headerData))

	log.Printf("- Body:\n%s\n", formatJSON(body))
	log.Println("========================")

	// 发送响应
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Webhook received successfully")

	// 记录处理时间
	log.Printf("Processed in %v\n", time.Since(start))
}

// 尝试格式化 JSON 内容
func formatJSON(body []byte) string {
	var out bytes.Buffer
	if err := json.Indent(&out, body, "", "  "); err != nil {
		return string(body) // 如果不是 JSON 就原样返回
	}
	return out.String()
}
