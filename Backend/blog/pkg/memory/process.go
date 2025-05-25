package memory

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
)

func (store *MessageStore) ProcessAuthorization(authorizationRequest *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	authResultChan := make(chan *pb.AuthorizationResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		response, exists := store.GetMessage("a" + authorizationRequest.Token)
		if exists {
			if authResponse, ok := response.(*pb.AuthorizationResponse); ok {
				authResultChan <- authResponse
			} else {
				errorChan <- fmt.Errorf("invalid type for authorization response")
			}
		} else {
			errorChan <- fmt.Errorf("authorization response not found")
		}
	}()

	select {
	case authorizationResponse := <-authResultChan:
		return authorizationResponse, nil
	case err := <-errorChan:
		return nil, err
	case <-time.After(time.Millisecond * 110):
		return nil, fmt.Errorf("authorization timeout")
	}
}

func (store *MessageStore) ProcessPushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) (*pb.GetUserFriendResponse, error) {
	friendResultChan := make(chan *pb.GetUserFriendResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		response, exists := store.GetMessage("f" + strconv.Itoa(int(getUserFriendRequest.Id)))
		if exists {
			if friendResponse, ok := response.(*pb.GetUserFriendResponse); ok {
				friendResultChan <- friendResponse
			} else {
				errorChan <- fmt.Errorf("invalid type for friend response")
			}
		} else {
			errorChan <- fmt.Errorf("friend response not found")
		}
	}()

	select {
	case friendResponse := <-friendResultChan:
		return friendResponse, nil
	case err := <-errorChan:
		return nil, err
	case <-time.After(time.Millisecond * 110):
		return nil, fmt.Errorf("push user friend timeout")
	}
}

func (store *MessageStore) ProcessPushUsers(getUsersRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	usersResultChan := make(chan *pb.GetUsersResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		key := "s"
		for _, id := range getUsersRequest.Ids {
			key += strconv.Itoa(int(id)) + ","
		}
		key = key[:len(key)-1]

		response, exists := store.GetMessage(key)
		if !exists {
			errorChan <- fmt.Errorf("user responses not found for ids: %v", getUsersRequest.Ids)
			return
		}
		if usersResponse, ok := response.(*pb.GetUsersResponse); ok {
			usersResultChan <- usersResponse
		} else {
			errorChan <- fmt.Errorf("invalid type for users response")
		}
	}()

	select {
	case usersResponse := <-usersResultChan:
		return usersResponse, nil
	case err := <-errorChan:
		return nil, err
	case <-time.After(time.Millisecond * 110):
		return nil, fmt.Errorf("push users timeout")
	}
}

func (store *MessageStore) ProcessPushUser(getUserRequest *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userResultChan := make(chan *pb.GetUserResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		response, exists := store.GetMessage("u" + strconv.Itoa(int(getUserRequest.Id)))
		if exists {
			if userResponse, ok := response.(*pb.GetUserResponse); ok {
				userResultChan <- userResponse
			} else {
				errorChan <- fmt.Errorf("invalid type for user response")
			}
		} else {
			errorChan <- fmt.Errorf("user response not found")
		}
	}()

	select {
	case userResponse := <-userResultChan:
		return userResponse, nil
	case err := <-errorChan:
		return nil, err
	case <-time.After(time.Millisecond * 110):
		return nil, fmt.Errorf("push user timeout")
	}
}
