package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/dom/user/internal/config"
	"github.com/labstack/echo"
)

type KafkaMessageRequest struct {
	Message string `json: "message"`
	Topic   string `json: "topic"`
}
type KafkaMessageResponse struct {
	Status string `json: "status"`
}

type KafkaHandler struct {
}

type KafkaHandlerParams struct {
}

func NewKafkaHandler(p *KafkaHandlerParams) *KafkaHandler {
	return &KafkaHandler{}
}

func (this KafkaHandler) Handle(c echo.Context) error {
	messageRequest := new(KafkaMessageRequest)

	if err := c.Bind(messageRequest); err != nil {
		return err
	}

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 5
	saramaConfig.Producer.Return.Successes = true
	kafkaConfig := config.DefaultKafkaConfig()
	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%s", kafkaConfig.Host, strconv.Itoa(kafkaConfig.Port))}, saramaConfig)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %s", err.Error())
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: messageRequest.Topic,
		Value: sarama.StringEncoder(messageRequest.Message),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error sending message to Kafka: %s", err.Error())
	}
	log.Printf("Message sent to partition %d at offset %d", partition, offset)

	response := &KafkaMessageResponse{
		Status: "Message sent",
	}

	return c.JSON(http.StatusOK, response)
}
