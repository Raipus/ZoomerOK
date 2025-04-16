package broker

import (
	"log"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"google.golang.org/protobuf/proto"
)

func (broker *RealBroker) Listen() {
	for {
		m, err := broker.reader.ReadMessage(broker.parent)
		if err != nil {
			log.Fatal("error while receiving message: %s", err.Error())
			continue
		}

		value := m.Value

		switch m.Key[0] {
		case 'u': // использует первый байт сообщения
			var response pb.GetUserResponse
			err = proto.Unmarshal(value, &response)
			if err != nil {
				log.Println("error: cannot deserialize GetUserRequest data")
				continue
			}
			log.Println(response)
		case 'f': // использует первый байт сообщения
			var friendResponse pb.GetUserFriendResponse
			err = proto.Unmarshal(value, &friendResponse)
			if err != nil {
				log.Println("error: cannot deserialize GetUserFriendRequest data")
				continue
			}
			log.Println(friendResponse)
		default:
			log.Println("unknown message type")
		}
	}
}
