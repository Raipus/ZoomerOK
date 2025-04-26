package broker

import (
	"log"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type BrokerInterface interface {
	PushUser(getUserRequest *pb.GetUserRequest) error
	PushUsers(getUsersRequest *pb.GetUsersRequest) error
	PushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) error
	Authorization(authorizationRequest *pb.AuthorizationRequest) error
	Listen()
}

var ProductionBrokerInterface BrokerInterface = &RealBroker{
	reader: initReader(),
	writer: initWriter(),
}

type RealBroker struct {
	reader *kafka.Reader
	writer *kafka.Writer
}

func initReader() *kafka.Reader {
	if gin.Mode() == gin.ReleaseMode {
		log.Println(config.Config.KafkaBrokerUrl)
		log.Println("Kafka reader initialized")
		kafkaConfig := kafka.ReaderConfig{
			Brokers:         []string{config.Config.KafkaBrokerUrl},
			Topic:           config.Config.KafkaReaderTopic,
			GroupID:         "account-service",
			MinBytes:        1,
			MaxBytes:        1024 * 1024,
			MaxWait:         1 * time.Second,
			ReadLagInterval: -1,
		}

		return kafka.NewReader(kafkaConfig)
	} else {
		return nil
	}
}

func initWriter() *kafka.Writer {
	if gin.Mode() == gin.ReleaseMode {
		log.Println("Kafka writer initialized")
		dialer := &kafka.Dialer{
			Timeout: time.Duration(10) * time.Second,
		}

		kafkaConfig := kafka.WriterConfig{
			Brokers:      []string{config.Config.KafkaBrokerUrl},
			Topic:        config.Config.KafkaWriterTopic,
			Balancer:     &kafka.LeastBytes{},
			Dialer:       dialer,
			WriteTimeout: 10 * time.Second,
			BatchSize:    1,
			BatchTimeout: 1 * time.Millisecond,
		}
		return kafka.NewWriter(kafkaConfig)
	} else {
		return nil
	}
}
