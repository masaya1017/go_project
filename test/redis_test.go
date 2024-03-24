package main

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	// Redisとの接続設定
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redisのアドレス
		Password: "",               // Redisのパスワード
		DB:       0,                // 使用するデータベース番号
	})

	// コンテキスト
	ctx := context.Background()

	// キーを設定
	err := rdb.Set(ctx, "mykey", "test", 0).Err()
	if err != nil {
		panic(err)
	}

	// キーを取得
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "test", val, "Retrieved value does not match expected value")

	// キーを削除する
	err = rdb.Del(ctx, "mykey").Err()
	if err != nil {
		panic(err)
	}

	// キーが削除されていることを確認する
	val, err = rdb.Get(ctx, "mykey").Result()

	// エラーチェック
	assert.Equal(t, redis.Nil, err, "Key still exists after deletion")
}
