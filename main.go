package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request") // リクエストを受けたら、ログ出力
		fmt.Fprintf(w, "Hello Docker!!") // Hello Docker!! とレスポンス
	})

	log.Println("start server")
	server := &http.Server{Addr: ":8080"} // 8080ポートでサーバーアプリケーションとして動作
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}