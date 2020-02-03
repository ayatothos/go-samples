package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/sample1", Sample1)
	http.HandleFunc("/sample2", Sample2)
	http.HandleFunc("/sample3", Sample3)
	http.ListenAndServe(":8080", nil)
}

// Sample1 ハンドラ
func Sample1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hoge", r.FormValue("hoge"))
	fmt.Println("foo", r.FormValue("foo"))
}

// Sample2 ハンドラ
func Sample2(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

// Sample3 ハンドラ
func Sample3(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Println("errorだよ")
	}

	mf := r.MultipartForm

	// 通常のリクエスト
	for k, v := range mf.Value {
		fmt.Printf("%v : %v\n", k, v)
	}

	// ファイル等バイナリのリクエスト
	for k, v := range mf.File {
		for kk, vv := range v {
			fmt.Printf("%v : %v : %v\n", k, kk, vv.Filename)
			fmt.Printf("%v : %v : %v\n", k, kk, vv.Header)
			fmt.Printf("%v : %v : %v\n", k, kk, vv.Size)
		}
	}
}
