// lib/parse_json.go
package lib

import (
	"encoding/json"
	"os"
)

// 構造体定義
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// JSONファイルをパースしてPerson構造体のリストにマッピングする関数
func ParseJSONFile(filename string) ([]Person, error) {
	// list形式のデータを取得する
	var people []Person
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		return people, err
	}
	defer file.Close()
	// JSONをデコード
	err = json.NewDecoder(file).Decode(&people)
	return people, err
}
