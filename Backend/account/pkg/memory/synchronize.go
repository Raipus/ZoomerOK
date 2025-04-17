package memory

import (
	"context"
	"sync"
)

var RedisContext = context.Background()
var RedisMu sync.Mutex
var RedisWaitGroup sync.WaitGroup
