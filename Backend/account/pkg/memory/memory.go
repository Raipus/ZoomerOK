package memory

import (
	"context"
	"log"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	GetUser(redisId int) RedisUser
	SetUser(redisUser RedisUser)
	DeleteUser(userId int)
	GetUsers(userIds []int) []RedisUser
	GetUserFriends(userId int) RedisUserFriend
	AddUserFriend(redisUserFriend RedisUserFriend)
	DeleteAllUserFriend(userId int)
	DeleteUserFriend(userId, userFriendId int)
	GetAuthorization(token string) RedisAuthorization
	SetAuthorization(redisAuthorization RedisAuthorization)
	DeleteAuthorization(token string)
}

var ProductionRedisInterface RedisInterface = &RealRedis{client: initClient()}

type RealRedis struct {
	client *redis.Client
}

func initClient() *redis.Client {
	if gin.Mode() == gin.ReleaseMode {
		log.Println(config.Config.RedisUrl)
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.Config.RedisUrl,
			Password: config.Config.RedisPassword,
			DB:       0,
		})

		// Проверка подключения
		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Ошибка подключения к Redis: %v", err)
		}

		log.Println("Redis Connected!")
		return rdb
	} else {
		return nil
	}
}
