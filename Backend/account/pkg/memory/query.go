package memory

import (
	"fmt"
	"log"
	"strconv"
)

type RedisUser struct {
	UserId int
	Login  string
	Name   string
	Image  string
}

type RedisUserFriend struct {
	UserId    int
	FriendIds []int
}

type RedisAuthorization struct {
	UserId         int
	Token          string
	Login          string
	Email          string
	ConfirmedEmail bool
}

func CompareRedisUser(redisUser1, redisUser2 RedisUser) bool {
	return redisUser1.UserId == redisUser2.UserId
}

func CompareRedisUserFriend(redisUserFriend1, redisUserFriend2 RedisUserFriend) bool {
	return redisUserFriend1.UserId == redisUserFriend2.UserId
}

func CompareRedisAuthorization(redisAuthorization1, redisAuthorization2 RedisAuthorization) bool {
	return redisAuthorization1.UserId == redisAuthorization2.UserId
}

func (r *RealRedis) GetUser(userId int) RedisUser {
	RedisMu.Lock()
	userData, err := r.client.HMGet(RedisContext, "user_"+string(userId), "login", "name", "image").Result()
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return RedisUser{}
	}
	RedisMu.Unlock()

	user := RedisUser{
		UserId: userId,
		Login:  userData[0].(string),
		Name:   userData[1].(string),
		Image:  userData[2].(string),
	}

	return user
}

func (r *RealRedis) SetUser(redisUser RedisUser) {
	RedisMu.Lock()
	err := r.client.HMSet(RedisContext, "user_"+string(redisUser.UserId), map[string]interface{}{
		"login": redisUser.Login,
		"name":  redisUser.Name,
		"image": redisUser.Image,
	}).Err()

	if err != nil {
		log.Fatalf("Cannot save user: %s", err)
	}

	RedisMu.Unlock()
}

func (r *RealRedis) DeleteUser(userId int) {
	RedisMu.Lock()
	err := r.client.Del(RedisContext, "user_"+string(userId)).Err()

	if err != nil {
		log.Fatalf("Cannot delete user: %s", err)
	}

	RedisMu.Unlock()
}

func (r *RealRedis) GetUsers(userIds []int) []RedisUser {
	RedisMu.Lock()
	if len(userIds) == 0 {
		return []RedisUser{}
	}
	var users []RedisUser
	for _, friendId := range userIds {
		user := r.GetUser(friendId)
		if user.UserId != 0 { // Check if user was successfully retrieved
			users = append(users, user)
		}
	}
	RedisMu.Unlock()

	return users
}

func (r *RealRedis) GetUserFriends(userId int) RedisUserFriend {
	RedisMu.Lock()
	friendIds, err := r.client.SMembers(RedisContext, "user_"+string(userId)+"_friends").Result()
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
	for _, friendId := range redisUserFriend.FriendIds {
		r.client.SAdd(RedisContext, "user_"+string(redisUserFriend.UserId)+"_friends", friendId)
	}
	RedisMu.Unlock()
}

func (r *RealRedis) DeleteAllUserFriend(userId int) {
	RedisMu.Lock()
	err := r.client.Del(RedisContext, "user_"+string(userId)+"_friends").Err()
	if err != nil {
		log.Fatalf("Cannot delete all user friends: %s", err)
	}
	RedisMu.Unlock()
}

func (r *RealRedis) DeleteUserFriend(userId, userFriendId int) {
	RedisMu.Lock()
	err := r.client.SRem(RedisContext, "user_"+string(userId)+"_friends", userFriendId).Err()
	if err != nil {
		log.Printf("Ошибка при удалении друга: %v", err)
	}
	RedisMu.Unlock()
}

func (r *RealRedis) GetAuthorization(token string) RedisAuthorization {
	log.Println(333)
	RedisMu.Lock()
	log.Println(324)
	userData, err := r.client.HMGet(RedisContext, "auth_"+token, "id", "login", "email", "confirmed_email").Result()
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return RedisAuthorization{}
	}
	RedisMu.Unlock()

	userIdString, ok := userData[0].(string)
	if !ok {
		log.Printf("Ошибка: UserId не является строкой.  Тип: %T, Значение: %v", userData[0], userData[0])
		userIdString = "-1"
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		log.Printf("Ошибка при преобразовании UserId в int: %v, UserIdString: %s", err, userIdString)
		userId = -1
	}

	var confirmedEmail bool = false
	strConfirmedEmail := userData[3].(string)
	log.Println(234)
	if strConfirmedEmail == "1" {
		confirmedEmail = true
	} else {
		confirmedEmail = false
	}

	authorization := RedisAuthorization{
		UserId:         userId,
		Token:          token,
		Login:          userData[1].(string),
		Email:          userData[2].(string),
		ConfirmedEmail: confirmedEmail,
	}

	return authorization
}

func (r *RealRedis) SetAuthorization(redisAuthorization RedisAuthorization) {
	RedisMu.Lock()
	err := r.client.HMSet(RedisContext, "auth_"+redisAuthorization.Token, map[string]interface{}{
		"id":              redisAuthorization.UserId,
		"login":           redisAuthorization.Login,
		"email":           redisAuthorization.Email,
		"confirmed_email": redisAuthorization.ConfirmedEmail,
	}).Err()

	if err != nil {
		log.Fatalf("Cannot save authorization: %s", err)
	}

	RedisMu.Unlock()
}

func (r *RealRedis) DeleteAuthorization(token string) {
	RedisMu.Lock()
	err := r.client.Del(RedisContext, "auth_"+token).Err()

	if err != nil {
		log.Fatalf("Cannot delete authorization: %s", err)
	}

	RedisMu.Unlock()
}
