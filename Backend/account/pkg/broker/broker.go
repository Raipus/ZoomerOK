package broker

import (
	"log"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type BrokerInterface interface {
	PushUser(getUserResponse *pb.GetUserResponse) error
	PushUsers(GetUsersResponse *pb.GetUsersResponse) error
	PushUserFriend(getUserFriendResponse *pb.GetUserFriendResponse) error
	Authorization(authorizationResponse *pb.AuthorizationResponse) error
	Listen()
}

var ProductionBrokerInterface BrokerInterface = &RealBroker{
	mem:    memory.ProductionRedisInterface,
	reader: initReader(),
	writer: initWriter(),
}

type RealBroker struct {
	mem    memory.RedisInterface
	reader *kafka.Reader
	writer *kafka.Writer
}

func initReader() *kafka.Reader {
	if gin.Mode() == gin.ReleaseMode {
		log.Println("Kafka reader initialized")
		log.Println(config.Config.KafkaBrokerUrl)
		kafkaConfig := kafka.ReaderConfig{
			Brokers:         []string{config.Config.KafkaBrokerUrl},
			Topic:           config.Config.KafkaReaderTopic,
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
