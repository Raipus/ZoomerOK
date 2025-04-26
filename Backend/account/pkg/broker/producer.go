package broker

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
)

func (broker *RealBroker) PushUser(getUserResponse *pb.GetUserResponse) error {
	log.Println("Push User")
	data, err := proto.Marshal(getUserResponse)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("u" + strconv.Itoa(int(getUserResponse.Id))),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

func (broker *RealBroker) PushUsers(getUsersResponse *pb.GetUsersResponse) error {
	log.Println("Push Users")
	data, err := proto.Marshal(getUsersResponse)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	var idStrings []string
	for _, id := range getUsersResponse.Ids {
		idStrings = append(idStrings, strconv.Itoa(int(id)))
	}
	idsString := strings.Join(idStrings, ",")
	message := kafka.Message{
		Key:   []byte("s" + idsString),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

func (broker *RealBroker) PushUserFriend(getUserFriendResponse *pb.GetUserFriendResponse) error {
	log.Println("Push User Friend")
	data, err := proto.Marshal(getUserFriendResponse)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("f" + strconv.Itoa(int(getUserFriendResponse.Id))),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

func (broker *RealBroker) Authorization(authorizationResponse *pb.AuthorizationResponse) error {
	log.Println("Authorization")
	data, err := proto.Marshal(authorizationResponse)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	message := kafka.Message{
		Key:   []byte("a" + authorizationResponse.Token),
		Value: data,
	}
	if err := broker.writer.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}
