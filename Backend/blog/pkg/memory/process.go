package memory

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
)

func (store *MessageStore) ProcessAuthorization(authorizationRequest *pb.AuthorizationRequest) (pb.AuthorizationResponse, error) {
	authResultChan := make(chan pb.AuthorizationResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		response, exists := store.GetMessage("a" + authorizationRequest.Token)
		if exists {
			authResultChan <- response.(pb.AuthorizationResponse)
		} else {
			errorChan <- fmt.Errorf("authorization response not found")
		}
	}()

	select {
	case authorizationResponse := <-authResultChan:
		return authorizationResponse, nil
	case err := <-errorChan:
		return pb.AuthorizationResponse{}, err
	case <-time.After(time.Millisecond * 110):
		return pb.AuthorizationResponse{}, fmt.Errorf("authorization timeout")
	}
}

func (store *MessageStore) ProcessPushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) (pb.GetUserFriendResponse, error) {
	friendResultChan := make(chan pb.GetUserFriendResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		response, exists := store.GetMessage("f" + strconv.Itoa(int(getUserFriendRequest.Id)))
		if exists {
			friendResultChan <- response.(pb.GetUserFriendResponse)
		} else {
			errorChan <- fmt.Errorf("friend response not found")
		}
	}()

	select {
	case friendResponse := <-friendResultChan:
		return friendResponse, nil
	case err := <-errorChan:
		return pb.GetUserFriendResponse{}, err
	case <-time.After(time.Millisecond * 110):
		return pb.GetUserFriendResponse{}, fmt.Errorf("push user friend timeout")
	}
}

func (store *MessageStore) ProcessPushUsers(getUsersRequest *pb.GetUsersRequest) (pb.GetUsersResponse, error) {
	usersResultChan := make(chan pb.GetUsersResponse)
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
		usersResultChan <- response.(pb.GetUsersResponse)
	}()

	select {
	case usersResponse := <-usersResultChan:
		return usersResponse, nil
	case err := <-errorChan:
		return pb.GetUsersResponse{}, err
	case <-time.After(time.Millisecond * 110):
		return pb.GetUsersResponse{}, fmt.Errorf("push users timeout")
	}
}

func (store *MessageStore) ProcessPushUser(getUserRequest *pb.GetUserRequest) (pb.GetUserResponse, error) {
	userResultChan := make(chan pb.GetUserResponse)
	errorChan := make(chan error)

	go func() {
		time.Sleep(time.Millisecond * 100)
		response, exists := store.GetMessage("u" + strconv.Itoa(int(getUserRequest.Id)))
		if exists {
			userResultChan <- response.(pb.GetUserResponse)
		} else {
			errorChan <- fmt.Errorf("user response not found")
		}
	}()

	select {
	case userResponse := <-userResultChan:
		return userResponse, nil
	case err := <-errorChan:
		return pb.GetUserResponse{}, err
	case <-time.After(time.Millisecond * 110):
		return pb.GetUserResponse{}, fmt.Errorf("push user timeout")
	}
}
