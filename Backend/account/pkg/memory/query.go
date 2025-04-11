package memory

import (
	"fmt"
	"log"
)

type RedisUser struct {
	UserId int
	Name   string
	Image  string
}

type RedisUserFriend struct {
	UserId    int
	FriendIds []int
}

func (r *RealRedis) GetUser(userId int) RedisUser {
	RedisMu.Lock()
	userData, err := r.client.HMGet(RedisContext, "user_"+string(userId), "name", "image").Result()
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return RedisUser{}
	}
	RedisMu.Unlock()

	user := RedisUser{
		UserId: userId,
		Name:   userData[0].(string),
		Image:  userData[1].(string),
	}

	return user
}

func (r *RealRedis) SetUser(redisUser RedisUser) {
	RedisMu.Lock()
	err := r.client.HMSet(RedisContext, "user_"+string(redisUser.UserId), map[string]interface{}{
		"name":  redisUser.Name,
		"image": redisUser.Image,
	}).Err()

	if err != nil {
		log.Fatalf("Cannot save user: %s", err)
	}

	RedisMu.Unlock()
}

func (r *RealRedis) DeleteUser(redisUser RedisUser) {
	RedisMu.Lock()
	err := r.client.Del(RedisContext, "user_"+string(redisUser.UserId)).Err()

	if err != nil {
		log.Fatalf("Cannot delete user: %s", err)
	}

	RedisMu.Unlock()
}

func (r *RealRedis) GetUserFriends(userId int) RedisUserFriend {
	RedisMu.Lock()
	friendIds, err := r.client.LRange(RedisContext, "user_"+string(userId)+"_friends", 0, -1).Result()
	if err != nil {
		log.Printf("Ошибка при получении друзей: %v", err)
		return RedisUserFriend{UserId: userId, FriendIds: []int{}}
	}
	RedisMu.Unlock()

	var friends []int
	for _, id := range friendIds {
		var friendId int
		if _, err := fmt.Sscanf(id, "%d", &friendId); err == nil {
			friends = append(friends, friendId)
		}
	}

	return RedisUserFriend{UserId: userId, FriendIds: friends}
}

func (r *RealRedis) AddUserFriend(redisUserFriend RedisUserFriend) {
	RedisMu.Lock()
	r.client.LPush(RedisContext, "user_"+string(redisUserFriend.UserId)+"_friends", redisUserFriend.FriendIds)
	RedisMu.Unlock()
}

func (r *RealRedis) DeleteUserFriend(redisUserFriend RedisUserFriend) {
	RedisMu.Lock()
	err := r.client.Del(RedisContext, "user_"+string(redisUserFriend.UserId)+"_friends").Err()

	if err != nil {
		log.Fatalf("Cannot delete user friend: %s", err)
	}
	RedisMu.Unlock()
}
