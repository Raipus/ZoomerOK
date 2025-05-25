package memory

import (
	"log"
	"sync"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
)

type MessageStoreInterface interface {
	SaveMessage(key string, message interface{}) error
	GetMessage(key string) (interface{}, bool)
	ProcessAuthorization(authorizationRequest *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error)
	ProcessPushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) (*pb.GetUserFriendResponse, error)
	ProcessPushUsers(getUsersRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error)
	ProcessPushUser(getUserRequest *pb.GetUserRequest) (*pb.GetUserResponse, error)
}

type MessageStore struct {
	messages map[string]interface{}
	mu       sync.RWMutex
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		messages: make(map[string]interface{}),
	}
}

func (store *MessageStore) SaveMessage(key string, message interface{}) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.messages[key] = message
	log.Println("Сохранил сообщение:", key)
	return nil
}

func (store *MessageStore) GetMessage(key string) (interface{}, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()
	msg, exists := store.messages[key]
	if exists {
		delete(store.messages, key)
	}
	log.Println("Достал сообщение:", key)
	return msg, exists
}

var ProductionMessageStore MessageStoreInterface = NewMessageStore()
