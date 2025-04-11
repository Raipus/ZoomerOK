package broker

import (
	"log"

	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
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
			var request pb.GetUserRequest
			err = proto.Unmarshal(value, &request)
			if err != nil {
				log.Println("error: cannot deserialize GetUserRequest data")
				continue
			}
			broker.HandleGetUserRequest(request)
		case 'f': // использует первый байт сообщения
			var friendRequest pb.GetUserFriendRequest
			err = proto.Unmarshal(value, &friendRequest)
			if err != nil {
				log.Println("error: cannot deserialize GetUserFriendRequest data")
				continue
			}
			broker.HandleGetUserFriendRequest(friendRequest)
		default:
			log.Println("unknown message type")
		}
	}
}

func (broker *RealBroker) HandleGetUserRequest(request pb.GetUserRequest) {
	user := broker.mem.GetUser(int(request.Id))
	response := &pb.GetUserResponse{
		Id:    int64(user.UserId),
		Name:  user.Name,
		Image: user.Image,
	}

	broker.PushUser(response)
}

func (broker *RealBroker) HandleGetUserFriendRequest(request pb.GetUserFriendRequest) {
	friends := broker.mem.GetUserFriends(int(request.Id))
	Int64FriendsIds := make([]int64, len(friends.FriendIds))

	for i, v := range friends.FriendIds {
		Int64FriendsIds[i] = int64(v)
	}

	response := &pb.GetUserFriendResponse{
		Id:  int64(friends.UserId),
		Ids: Int64FriendsIds,
	}
	broker.PushUserFriend(response)
}
