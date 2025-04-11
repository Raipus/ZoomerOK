package broker

import (
	"context"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/segmentio/kafka-go"
)

type BrokerInterface interface {
	PushUser(redisUser *pb.GetUserResponse) error
	PushUserFriend(redisUserFriend *pb.GetUserFriendResponse) error
	Listen()
}

var ProductionBrokerInterface BrokerInterface = &RealBroker{
	parent: context.Background(),
	mem:    memory.ProductionRedisInterface,
	reader: initReader(),
	writer: initWriter(),
}

type RealBroker struct {
	parent context.Context
	mem    memory.RedisInterface
	reader *kafka.Reader
	writer *kafka.Writer
}

func initReader() *kafka.Reader {
	kafkaConfig := kafka.ReaderConfig{
		Brokers:         []string{config.Config.KafkaBrokerUrl},
		Topic:           config.Config.KafkaAccountBlogTopic,
		MinBytes:        10e3,
		MaxBytes:        57671680,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}

	return kafka.NewReader(kafkaConfig)
}

func initWriter() *kafka.Writer {
	dialer := &kafka.Dialer{
		Timeout: time.Duration(10) * time.Second,
	}

	kafkaConfig := kafka.WriterConfig{
		Brokers:      []string{config.Config.KafkaBrokerUrl},
		Topic:        config.Config.KafkaAccountBlogTopic,
		Balancer:     &kafka.LeastBytes{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
	}
	return kafka.NewWriter(kafkaConfig)
}
