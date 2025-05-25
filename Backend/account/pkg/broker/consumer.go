package broker

import (
	"context"
	"log"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"google.golang.org/protobuf/proto"
)

func (broker *RealBroker) Listen() {
	log.Println("Start listening broker")
	for {
		time.Sleep(time.Millisecond)
		m, err := broker.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}

		value := m.Value

		switch m.Key[0] {
		case 'a': // использует первый байт сообщения
			var request pb.AuthorizationRequest
			err = proto.Unmarshal(value, &request)
			if err != nil {
				log.Println("error: cannot deserialize AuthorizationRequest data")
				continue
			}
			broker.HandleAuthorizationRequest(&request)
		case 'u': // использует первый байт сообщения
			var request pb.GetUserRequest
			err = proto.Unmarshal(value, &request)
			if err != nil {
				log.Println("error: cannot deserialize GetUserRequest data")
				continue
			}
			broker.HandleGetUserRequest(&request)
		case 's':
			var request pb.GetUsersRequest
			err = proto.Unmarshal(value, &request)
			if err != nil {
				log.Println("error: cannot deserialize GetUsersResponse data")
				continue
			}
			broker.HandleGetUsersRequest(&request)
		case 'f': // использует первый байт сообщения
			var friendRequest pb.GetUserFriendRequest
			err = proto.Unmarshal(value, &friendRequest)
			if err != nil {
				log.Println("error: cannot deserialize GetUserFriendRequest data")
				continue
			}
			broker.HandleGetUserFriendRequest(&friendRequest)
		default:
			log.Println("unknown message type")
		}
	}
}

func (broker *RealBroker) HandleAuthorizationRequest(request *pb.AuthorizationRequest) {
	authorization := broker.mem.GetAuthorization(request.Token)

	var response *pb.AuthorizationResponse
	if memory.CompareRedisAuthorization(authorization, memory.RedisAuthorization{}) {
		response = &pb.AuthorizationResponse{
			Id:             int64(0),
			Login:          "",
			Email:          "",
			Token:          "",
			ConfirmedEmail: false,
		}
	} else {
		response = &pb.AuthorizationResponse{
			Id:             int64(authorization.UserId),
			Login:          authorization.Login,
			Email:          authorization.Email,
			Token:          authorization.Token,
			ConfirmedEmail: authorization.ConfirmedEmail,
		}
	}

	log.Println(response)
	err := broker.Authorization(response)
	if err != nil {
		log.Printf("Ошибка при отправки сообщения: %v", err)
		return
	}
}

func (broker *RealBroker) HandleGetUserRequest(request *pb.GetUserRequest) {
	user := broker.mem.GetUser(int(request.Id))

	var response *pb.GetUserResponse
	if memory.CompareRedisUser(user, memory.RedisUser{}) {
		response = &pb.GetUserResponse{
			Id:    int64(0),
			Name:  "",
			Login: "",
			Image: "",
		}
	} else {
		response = &pb.GetUserResponse{
			Id:    int64(user.UserId),
			Name:  user.Name,
			Login: user.Login,
			Image: user.Image,
		}
	}

	log.Println(response)
	err := broker.PushUser(response)
	if err != nil {
		log.Printf("Ошибка при отправки сообщения: %v", err)
		return
	}
}

func (broker *RealBroker) HandleGetUsersRequest(request *pb.GetUsersRequest) {
	IntUsersIds := make([]int, len(request.Ids))

	for i, v := range request.Ids {
		IntUsersIds[i] = int(v)
	}
	redisUsers := broker.mem.GetUsers(IntUsersIds)

	var response *pb.GetUsersResponse
	if len(redisUsers) == 0 {
		response = &pb.GetUsersResponse{
			Ids:   []int64{},
			Users: []*pb.GetUserResponse{},
		}
	} else {
		users := make([]*pb.GetUserResponse, 0, len(redisUsers))
		for _, redisUser := range redisUsers {
			users = append(users, &pb.GetUserResponse{
				Image: redisUser.Image,
				Name:  redisUser.Name,
				Login: redisUser.Login,
				Id:    int64(redisUser.UserId),
			})
		}
		response = &pb.GetUsersResponse{
			Ids:   request.Ids,
			Users: users,
		}
	}

	log.Println(response)
	err := broker.PushUsers(response)
	if err != nil {
		log.Printf("Ошибка при отправки сообщения: %v", err)
		return
	}
}

func (broker *RealBroker) HandleGetUserFriendRequest(request *pb.GetUserFriendRequest) {
	friends := broker.mem.GetUserFriends(int(request.Id))

	var response *pb.GetUserFriendResponse
	if memory.CompareRedisUserFriend(friends, memory.RedisUserFriend{}) {
		response = &pb.GetUserFriendResponse{
			Id:  int64(0),
			Ids: []int64{},
		}
	} else {
		Int64FriendsIds := make([]int64, len(friends.FriendIds))

		for i, v := range friends.FriendIds {
			Int64FriendsIds[i] = int64(v)
		}

		response = &pb.GetUserFriendResponse{
			Id:  int64(friends.UserId),
			Ids: Int64FriendsIds,
		}
	}

	log.Println(response)
	err := broker.PushUserFriend(response)
	if err != nil {
		log.Printf("Ошибка при отправки сообщения: %v", err)
		return
	}
}
