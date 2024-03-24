package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestConnectSqlite(t *testing.T) {
	// SQLiteデータベースに接続
	db, err := sql.Open("sqlite3", "../example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// タイムアウト時間を設定
	timeout := 5 * time.Second
	db.SetConnMaxLifetime(timeout)

	// データベースへの接続が成功したことを確認
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("SQLiteデータベースにアクセス成功しました:", "./example.db")
}

func TestGorm(t *testing.T) {
	// gormによる更新処理の実装
	type User struct {
		ID   uint
		Name string
		Age  uint
	}

	// SQLiteデータベースに接続
	db, err := gorm.Open(sqlite.Open("example.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// データベースの自動マイグレーション
	db.AutoMigrate(&User{})

	// ユーザーの作成
	user := User{Name: "Alice", Age: 30}
	db.Create(&user)

	// ユーザーの取得
	var result User
	db.Where(User{Name: "Alice"}).First(&result)
	log.Printf("Found user: %+v", result)

	// ユーザーの更新
	db.Model(&result).Update("Age", 31)

	log.Printf("Updated user: %+v", result)

	// ユーザーの削除
	db.Delete(&result)

}
