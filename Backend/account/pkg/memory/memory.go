package memory

import (
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	GetUser(redisId int) RedisUser
	SetUser(redisUser RedisUser)
	DeleteUser(redisUser RedisUser)
	GetUserFriends(userId int) RedisUserFriend
	AddUserFriend(redisUserFriend RedisUserFriend)
	DeleteUserFriend(redisUserFriend RedisUserFriend)
}

var ProductionRedisInterface RedisInterface = &RealRedis{client: initClient()}

type RealRedis struct {
	client *redis.Client
}

func initClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisUrl,
		Password: config.Config.RedisPassword,
		DB:       0,
	})

	return rdb
}
