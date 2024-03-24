package main

import (
	"go_project/lib"
	"testing"
)

// テストケース
func TestParseJSONFile(t *testing.T) {
	t.Skip("This test is skipped intentionally")
	// テスト用のJSONファイル
	filename := "../data/sample.json"
	//検索条件
	searchName := "Alice"

	// JSONファイルをパースして構造体にマッピング
	result, err := lib.ParseJSONFile(filename)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// マッピングされた結果をログに出力
	for i, person := range result {
		if person.Name == searchName {
			t.Logf("----------------")
			t.Logf("Person %d:", i+1)
			t.Logf("  Name: %s", person.Name)
			t.Logf("  Age: %d", person.Age)
			t.Logf("  Address: %s", person.Address)
		}
	}
	t.Logf("----------------")
}
