package memory

import (
	"sync"
)

type MessageQueue interface {
	Update(msg interface{})
	GetLastMessage() interface{}
}

type LastMessageQueue struct {
	mu       sync.Mutex
	messages []interface{} // предполагается, что Message - это общий тип для всех сообщений
}

func (q *LastMessageQueue) Update(msg interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.messages = append(q.messages, msg)
}

func (q *LastMessageQueue) GetLastMessage() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.messages) == 0 {
		return nil
	}

	lastMessage := q.messages[len(q.messages)-1]
	q.messages = q.messages[:len(q.messages)-1]
	return lastMessage
}

var ProductionLastMessageQueue MessageQueue = &LastMessageQueue{}
