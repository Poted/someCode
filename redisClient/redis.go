package redisClient

import (
	"codewars/session"
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
)

// docker run --hostname=277256d947e9 --mac-address=02:42:ac:11:00:02 --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOSU_VERSION=1.17 --env=REDIS_VERSION=7.4.0 --env=REDIS_DOWNLOAD_URL=http://download.redis.io/releases/redis-7.4.0.tar.gz --env=REDIS_DOWNLOAD_SHA=57b47c2c6682636d697dbf5d66d8d495b4e653afc9cd32b7adf9da3e433b8aaf --volume=/data --workdir=/data --runtime=runc -p 6379:6379 -d redis:latest

var ctx = context.Background()

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "",
		// DB: 0,
	})

	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}
	fmt.Println(status)

	return rdb
}

type RedisResponse struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func (s RedisResponse) ConvertToRedis() []RedisResponse {

	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return nil
	}

	var tags []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		// value := v.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		tags = append(tags, jsonTag)
	}

	fmt.Printf("tags: %v\n", tags)

	return nil
}

func StoreSession(rdb *redis.Client, sessionID uuid.UUID, sessionData *session.UserSession) error {
	err := rdb.Set(ctx, sessionID.String(), sessionData, 0).Err()
	return err
}

func GetSession(rdb *redis.Client, sessionID uuid.UUID) (string, error) {
	val, err := rdb.Get(ctx, sessionID.String()).Result()
	return val, err
}

func StoreSessionWithExpiration(rdb *redis.Client, sessionID uuid.UUID, sessionData session.UserSession, expiration time.Duration) error {
	err := rdb.Set(ctx, sessionID.String(), sessionData, expiration).Err()
	return err
}

func DeleteSession(rdb *redis.Client, sessionID uuid.UUID) error {
	err := rdb.Del(ctx, sessionID.String()).Err()
	return err
}
