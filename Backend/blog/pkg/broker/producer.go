package broker

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
)

func (broker *RealBroker) PushUser(getUserRequest *pb.GetUserRequest) error {
	log.Println("Push User")
	data, err := proto.Marshal(getUserRequest)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("u" + strconv.Itoa(int(getUserRequest.Id))),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

func (broker *RealBroker) PushUsers(getUsersRequest *pb.GetUsersRequest) error {
	log.Println("Push Users")
	var idStrings []string
	for _, id := range getUsersRequest.Ids {
		idStrings = append(idStrings, strconv.Itoa(int(id)))
	}
	idsString := strings.Join(idStrings, ",")
	data, err := proto.Marshal(getUsersRequest)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("s" + idsString),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

func (broker *RealBroker) PushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) error {
	log.Println("Push User Friend")
	data, err := proto.Marshal(getUserFriendRequest)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("f" + strconv.Itoa(int(getUserFriendRequest.Id))),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

func (broker *RealBroker) Authorization(authorizationRequest *pb.AuthorizationRequest) error {
	log.Println("Push Authorization")
	data, err := proto.Marshal(authorizationRequest)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("a" + authorizationRequest.Token),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}
