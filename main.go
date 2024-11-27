package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "테스트 엔드포인트입니다.")
	})

	http.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "아이템 이름은 %s입니다.", r.URL.Query().Get("name"))
	})

	fmt.Println("서버를 포트 8080에서 시작했습니다.")
	http.ListenAndServe(":8080", nil)
}
