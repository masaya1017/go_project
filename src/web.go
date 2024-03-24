package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	// パスパラメータを受け取り、それを表示するエンドポイント
	router.HandleFunc("/hello/{name}", NameHandler)
	router.HandleFunc("/products", ProductsHandler)
	router.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	// 挨拶対象の名前を取得する
	vars := mux.Vars(r)
	name := vars["name"]
	// クエリパラメータを取得する(挨拶のメッセージを取得する)
	queryParams := r.URL.Query()
	greeting := queryParams.Get("greeting")

	if name != "" {
		if greeting != "" {
			fmt.Fprintf(w, "%s, %s!", greeting, name)
		} else {
			fmt.Fprintf(w, "Hello, %s!", name)
		}
	} else {
		// nameが空の場合はデフォルトのメッセージを表示する
		fmt.Fprintf(w, "Hello, Guest!")
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products Page2")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Articles Page")
}
