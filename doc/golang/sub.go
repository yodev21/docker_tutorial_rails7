// サーバーアプケーションとして動作し、以下の挙動をします。
// ●どんなHTTPリクエストに対しても「Hello Docker!!」とレスポンスする
// ●8080ポートでサーバアプリケーションとして動作する
// ●クライアントからリクエストを受けた際は、received requestのログを標準出力に表示する
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		log.Println("received request")
		fmt.Fprintf(w, "Hello sub!!")
	})

	log.Println("start server")
	server := &http.Server{Addr: ":8090"}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}