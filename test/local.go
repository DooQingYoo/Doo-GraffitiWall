// 用于从本地的redis缓存里查找是否有相应的文件
// 以及将文件存到缓存中
package main

import (
	"github.com/go-redis/redis"
	"log"
	"myProject/test/config"
	"time"
)

const (
	markdownExpireTime  = time.Hour
	imageExpireTime     = 24 * time.Hour
	direcotryExpireTime = 30 * time.Minute
)

var redisClt *redis.Client

/**
 * 从缓存里找文件
 */
func findLocal(path string) (string, error) {
	result, err := redisClt.Get(path).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

/**
 * 把文件存到缓存中
 */
func saveLocal(path string, content string, duration time.Duration) {
	setResult := redisClt.Set(path, content, duration)
	if setResult.Err() != nil {
		log.Fatal(setResult.Err().Error())
	}
}

func init() {
	redisClt = redis.NewClient(config.Config.Redis)
	err := redisClt.Ping().Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}
