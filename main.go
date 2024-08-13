package main

import (
	"codewars/redisClient"
	"codewars/session"
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {

	rc := redisClient.ConnectRedis()

	sess := session.StartSession()

	sessID, err := uuid.NewV7()
	if err != nil {
		fmt.Printf("uuid err: %v\n", err)
	}

	err = redisClient.StoreSession(rc, sessID, sess)
	if err != nil {
		fmt.Printf("save err: %v\n", err)
	}

	rs := redisClient.RedisResponse{
		Key:   "some",
		Value: 200,
	}.ConvertToRedis()

	fmt.Printf("rs: %v\n", rs)

	xd, err := redisClient.GetSession(rc, sessID)
	if err != nil {
		fmt.Printf("get err: %v\n", err)
	}

	fmt.Printf("xd: %v\n", xd)

}
