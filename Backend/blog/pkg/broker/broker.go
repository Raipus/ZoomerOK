package broker

import (
	"context"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/config"
	"github.com/segmentio/kafka-go"
)

type BrokerInterface interface {
	PushGetUser(redisUser *pb.GetUserRequest) error
	PushGetUserFriend(redisUserFriend *pb.GetUserFriendRequest) error
	Listen()
}

var ProductionBrokerInterface BrokerInterface = &RealBroker{
	parent: context.Background(),
	reader: initReader(),
	writer: initWriter(),
}

type RealBroker struct {
	parent context.Context
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
