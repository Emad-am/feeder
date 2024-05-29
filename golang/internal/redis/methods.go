package redis

import (
	"encoding/json"
	"fmt"

	"github.com/Emad-am/feeder/internal/context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.GetContext()

func Lpush(key string, value interface{}) int64 {
	v, e := json.Marshal(value)
	if e != nil {
		fmt.Println(e)
	}

	res, err := Rdb.LPush(*ctx, key, v).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return res
}

func Rpop(key string) map[string]string {
	res, err := Rdb.RPop(*ctx, key).Result()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var dat map[string]string

	_ = json.Unmarshal([]byte(res), &dat)

	return dat
}

func Publish(channel string, payload interface{}) {
	err := Rdb.Publish(*ctx, channel, payload).Err()
	if err != nil {
		panic(err)
	}
}

func Subscribe(channel string) *redis.PubSub {
	pubsub := Rdb.Subscribe(*ctx, channel)
	return pubsub
}

func Hset(key string, values ...interface{}) {
	err := Rdb.HSet(*ctx, key, values...).Err()
	if err != nil {
		panic(err)
	}
}
