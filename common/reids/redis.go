package reids

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

var (
	once        sync.Once
	redisClient *redis.Client
	LockTimeOut = 3 * time.Second
	redisOption = new(redis.Options)
)

func init() {
	redisInit()
}

func redisInit() {
	once.Do(func() {
		redisOption.Addr = "127.0.0.1:6379"
		redisOption.DB = 0
		redisOption.Password = ""
		redisClient = redis.NewClient(redisOption)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for i := 0; i < 10; i++ {
			ping, err := redisClient.Ping(ctx).Result()
			if err != nil {
				log.Println(ping, err)
			} else {
				break
			}
			time.Sleep(1 * time.Second)
		}
		log.Println("Redis init success...")
	})
}

// Get 从Redis获取string值
func Get(key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticket, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Println("从Redis获取", key, "失败:", err)
		return ""
	}
	return ticket
}

// Set 设置string值到Redis
func Set(value, key string, timestamp time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisClient.Set(ctx, key, value, timestamp).Result()
	if err != nil {
		log.Println("保存", key, "到Redis失败:", err)
		return
	}
}

// Del 从Redis删除Key
func Del(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisClient.Del(ctx, key).Result()
	if err != nil {
		log.Println("删除", key, "失败:", err)
		return
	}
}

// HSet 保存到hash
func HSet(key, field string, value interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisClient.HSet(ctx, key, field, value).Result()
	if err != nil {
		log.Println("保存hash", key, field, "到Redis失败:", err)
		return
	}
}

// HSetWithError 保存到Hash并返回错误（如果有）
func HSetWithError(key, field string, value interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return redisClient.HSet(ctx, key, field, value).Result()
}

// HMSet 保存到hash
func HMSet(key string, field map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisClient.HMSet(ctx, key, field).Result()
	if err != nil {
		log.Println("保存hash", key, field, "到Redis失败:", err)
		return
	}
}

// HDel 从hash中删除field
func HDel(key string, fields ...string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := redisClient.HDel(ctx, key, fields...).Result()
	if err != nil {
		log.Println("从hash删除", key, fields, "失败:", err)
		return
	}
	log.Println("删除字段成功,字段:", fields, "成功删除条数:", rs)
}

// HGet 从hash获取值
func HGet(key, field string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := redisClient.HGet(ctx, key, field).Result()
	if err != nil {
		log.Println("从Redis获取hash", key, field, "失败:", err)
		return ""
	}
	return rs
}

// HGetAll 从hash获取全部值
func HGetAll(key string) map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		log.Println("从Redis获取hash", key, "失败:", err)
		return nil
	}
	return rs
}

// Lock 分布式锁 获得锁
func Lock(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := redisClient.SetNX(ctx, key, true, LockTimeOut).Result()
	if err != nil {
		log.Println("从Redis获取锁", key, "错误:", err)
		return false
	}
	return rs
}

// Unlock 分布式锁 释放锁
func Unlock(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := redisClient.Del(ctx, key).Result()
	if err != nil || rs == 0 {
		log.Println("释放分布式锁", key, "失败:", err)
		return false
	}
	return true
}
