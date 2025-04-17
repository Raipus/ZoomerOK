package broker

import (
	"log"
	"strconv"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
)

func (broker *RealBroker) PushUser(redisUser *pb.GetUserResponse) error {
	data, err := proto.Marshal(redisUser)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("u" + strconv.Itoa(int(redisUser.Id))),
		Value: data,
	}
	return broker.writer.WriteMessages(broker.parent, message)
}

func (broker *RealBroker) PushUserFriend(redisUserFriend *pb.GetUserFriendResponse) error {
	data, err := proto.Marshal(redisUserFriend)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("f" + strconv.Itoa(int(redisUserFriend.Id))),
		Value: data,
	}
	return broker.writer.WriteMessages(broker.parent, message)
}
