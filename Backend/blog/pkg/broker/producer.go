package broker

import (
	"strconv"

	"github.com/segmentio/kafka-go"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
)

func (broker *RealBroker) PushGetUser(redisUser *pb.GetUserRequest) error {
	message := kafka.Message{
		Key: []byte("u" + strconv.Itoa(int(redisUser.Id))),
	}
	return broker.writer.WriteMessages(broker.parent, message)
}

func (broker *RealBroker) PushGetUserFriend(redisUserFriend *pb.GetUserFriendRequest) error {
	message := kafka.Message{
		Key: []byte("f" + strconv.Itoa(int(redisUserFriend.Id))),
	}
	return broker.writer.WriteMessages(broker.parent, message)
}
