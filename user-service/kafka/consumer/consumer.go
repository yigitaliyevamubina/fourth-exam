package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "exam/user-service/genproto/user_service"
	"exam/user-service/pkg/logger"
	"exam/user-service/service/service"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	ConsumeMessages(handler func(message []byte, s *service.UserService)) error
	Close() error
}

type Consumer struct {
	reader *kafka.Reader
	s      *service.UserService
}

func NewKafkaConsumer(brokers []string, topic string, groupID string, s *service.UserService) (KafkaConsumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return &Consumer{reader: reader, s: s}, nil
}

func (c *Consumer) ConsumeMessages(handler func(message []byte, s *service.UserService)) error {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		handler(msg.Value, c.s)

		err = c.reader.CommitMessages(context.Background(), msg)
		if err != nil {
			return err
		}
	}

}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func ConsumeHandler(message []byte, s *service.UserService) {
	l := logger.New("", "")
	var user pb.User
	if err := json.Unmarshal(message, &user); err != nil {
		log.Fatal("cannot unmarshal json")
		return
	}
	user.Id = uuid.New().String()
	exists, err := s.CheckField(context.Background(), &pb.CheckFieldReq{Field: "email", Value: user.Email})
	if err != nil {
		l.Error("error while if user exists", logger.Error(err))
	}
	fmt.Println(exists.Status)
	if !exists.Status {
		resp, err := s.Create(context.Background(), &user)
		if err != nil {
			fmt.Println(err, "=====")
			l.Error("error while creating user via kafka")
			return
		}

		pp.Println("Successfully inserted user via kafka")
		pp.Println(resp)
	}
}
