package memory

import (
	"log"
	"sync"

	"encoding/base64"

	"github.com/Raipus/ZoomerOK/account/pkg/security"
)

type RedisCommand struct {
	Action string
	Params interface{}
}

var RedisCommandQueue = make(chan RedisCommand, 100)

type RedisAction interface {
	Execute() error
}

type AddUserAction struct {
	UserId int
	Name   string
	Image  []byte
}

func (a *AddUserAction) Execute(r RedisInterface) error {
	resizedImage, err := security.ResizeImage(a.Image)
	if err != nil {
		return err
	}

	base64Image := base64.StdEncoding.EncodeToString(resizedImage)
	redisUser := RedisUser{
		UserId: a.UserId,
		Name:   a.Name,
		Image:  base64Image,
	}

	r.SetUser(redisUser)

	return nil
}

func ProcessCommands(wg *sync.WaitGroup, r RedisInterface) {
	defer wg.Done()

	for command := range RedisCommandQueue {
		switch command.Action {
		case "resizeImage":
			action, ok := command.Params.(AddUserAction)
			if ok {
				if err := action.Execute(r); err != nil {
					log.Println(err)
				}
			}
		case "addFriend":
			action, ok := command.Params.(AddFriendAction)
			if ok {
				if err := action.Execute(r); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

type AddFriendAction struct {
	UserId    int
	FriendIds []int
}

func (a *AddFriendAction) Execute(r RedisInterface) error {
	redisUserFriend := RedisUserFriend{
		UserId:    a.UserId,
		FriendIds: a.FriendIds,
	}
	r.AddUserFriend(redisUserFriend)

	return nil
}
