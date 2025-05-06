package broker

import (
	"context"
	"log"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/memory"
	"google.golang.org/protobuf/proto"
)

func (broker *RealBroker) Listen() {
	log.Println("Start listening broker")
	for {
		time.Sleep(time.Millisecond)
		m, err := broker.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error()) // Используем Printf, а не Fatal
			break                                                        // Важно: выходим из цикла при ошибке чтения
		}

		value := m.Value

		switch m.Key[0] {
		case 'a': // использует первый байт сообщения
			var response pb.AuthorizationResponse
			err = proto.Unmarshal(value, &response)
			if err != nil {
				log.Println("error: cannot deserialize AuthorizationResponse data")
				continue
			}
			broker.HandleAuthorizationResponse(string(m.Key[:]), response)
		case 's': // использует первый байт сообщения
			var response pb.GetUsersResponse
			err = proto.Unmarshal(value, &response)
			if err != nil {
				log.Println("error: cannot deserialize GetUsersResponse data")
				continue
			}
			broker.HandleGetUsersResponse(string(m.Key[:]), response)
		case 'u': // использует первый байт сообщения
			var response pb.GetUserResponse
			err = proto.Unmarshal(value, &response)
			if err != nil {
				log.Println("error: cannot deserialize GetUserResponse data")
				continue
			}
			broker.HandleGetUserResponse(string(m.Key[:]), response)
		case 'f': // использует первый байт сообщения
			var friendResponse pb.GetUserFriendResponse
			err = proto.Unmarshal(value, &friendResponse)
			if err != nil {
				log.Println("error: cannot deserialize GetUserFriendResponse data")
				continue
			}
			broker.HandleGetUserFriendResponse(string(m.Key[:]), friendResponse)
		default:
			log.Println("unknown message type")
		}
	}
}

func (broker *RealBroker) HandleAuthorizationResponse(key string, response pb.AuthorizationResponse) {
	memory.ProductionMessageStore.SaveMessage(key, response)
}

func (broker *RealBroker) HandleGetUserResponse(key string, response pb.GetUserResponse) {
	memory.ProductionMessageStore.SaveMessage(key, response)
}

func (broker *RealBroker) HandleGetUsersResponse(key string, response pb.GetUsersResponse) {
	memory.ProductionMessageStore.SaveMessage(key, response)
}

func (broker *RealBroker) HandleGetUserFriendResponse(key string, response pb.GetUserFriendResponse) {
	memory.ProductionMessageStore.SaveMessage(key, response)
}
