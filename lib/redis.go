package lib

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// JSONデータをRedisに登録する関数
func SetJSONDataToRedis(ctx context.Context, rdb *redis.Client, key string, data interface{}) error {
	// JSONデータを文字列に変換
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Redisに文字列として登録
	err = rdb.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// RedisからJSONデータを取得する関数
func GetJSONDataFromRedis(ctx context.Context, rdb *redis.Client, key string, data interface{}) error {
	// Redisから文字列を取得
	jsonString, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	// 取得した文字列をJSONデコード
	err = json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return err
	}

	return nil
}

// Redisからデータを削除する関数
func DeleteDataFromRedis(ctx context.Context, rdb *redis.Client, key string) error {
	// キーを削除
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

// JSONデータをRedisに更新する関数
func UpdateJSONDataToRedis(ctx context.Context, rdb *redis.Client, key string, data interface{}) error {
	// JSONデータを文字列に変換
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Redisにキーが存在するかチェック
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return err
	}

	// キーが存在しない場合はエラーを返す
	if exists == 0 {
		return fmt.Errorf("Key does not exist")
	}

	// Redisに文字列として登録
	err = rdb.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
