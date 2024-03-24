package main

import (
	"fmt"
	"go_project/lib"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	// パスパラメータを受け取り、それを表示するエンドポイント
	router.HandleFunc("/hello/{name}", NameHandler)
	router.HandleFunc("/json", SearchHandler)
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

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// テスト用のJSONファイル
	filename := "data/sample.json"
	found := false
	// クエリパラメータを取得する(挨拶のメッセージを取得する)
	queryParams := r.URL.Query()
	searchName := queryParams.Get("name")

	// JSONファイルをパースして構造体にマッピング
	result, err := lib.ParseJSONFile(filename)
	if err != nil {
		fmt.Fprintf(w, "Unexpected error: %v", err) // エラーメッセージをクライアントに返す
		return
	}
	// マッピングされた結果をログに出力
	for _, person := range result {
		if person.Name == searchName {
			found = true
			fmt.Fprintf(w, "Age: %v", person.Age)
			fmt.Fprintf(w, "Address: %v", person.Address)
		}
	}
	// 検索結果が見つからなかった場合のメッセージをクライアントに返す
	if !found {
		fmt.Fprintf(w, "Person with name '%s' not found\n", searchName)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Articles Page")
}
