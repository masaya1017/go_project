package main

import (
	"context"
	"go_project/lib"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedis1(t *testing.T) {
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

func TestRedis2(t *testing.T) {
	// Redisとの接続設定
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redisのアドレス
		Password: "",               // Redisのパスワード
		DB:       0,                // 使用するデータベース番号
	})

	// コンテキスト
	ctx := context.Background()

	// JSONデータ
	userData := lib.User{
		Name:    "Alice",
		Age:     25,
		Address: "456 Elm St",
	}

	// JSONデータをRedisに登録
	err := lib.SetJSONDataToRedis(ctx, rdb, "user:id:100", userData)
	if err != nil {
		panic(err)
	}

	// RedisからJSONデータを取得
	var retrievedUser lib.User
	err = lib.GetJSONDataFromRedis(ctx, rdb, "user:id:100", &retrievedUser)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Alice", retrievedUser.Name, "Retrieved value does not match expected value")

	// 更新用のデータ
	updatedData := lib.User{
		Name:    "Bob",
		Age:     30,
		Address: "789 Oak St",
	}

	// Redisにデータを更新
	err = lib.UpdateJSONDataToRedis(ctx, rdb, "user:id:100", updatedData)
	if err != nil {
		panic(err)
	}

	// 更新したデータを再度取得
	var updatedUser lib.User
	err = lib.GetJSONDataFromRedis(ctx, rdb, "user:id:100", &updatedUser)
	if err != nil {
		panic(err)
	}

	// 更新されたデータを表示
	assert.Equal(t, "Bob", updatedUser.Name, "Retrieved value does not match expected value")

	// データを削除
	err = lib.DeleteDataFromRedis(ctx, rdb, "user:id:100")
	if err != nil {
		panic(err)
	}

	// データが削除されたことを確認
	var deletedUser lib.User
	err = lib.GetJSONDataFromRedis(ctx, rdb, "user:id:100", &deletedUser)
	assert.Error(t, err, "Data was not deleted")
}
