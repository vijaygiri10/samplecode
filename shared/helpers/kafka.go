package helpers

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/Shopify/sarama"

	"jetsend_opens/shared/log"
)

//GetProducer ...
func GetProducer(brokers []string, retryMax int, errorStatus bool) (*sarama.AsyncProducer, error) {
	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = errorStatus
	configs.Producer.Retry.Max = retryMax
	configs.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewAsyncProducer(brokers, configs)
	return &producer, err
}

//GetConsumerGroup ...
func GetConsumerGroup(brokers []string, consumerGroupID string, retryMax int, errorStatus bool) (*sarama.Client, *sarama.ConsumerGroup, error) {
	ctx := context.Background()
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	config.Consumer.Return.Errors = errorStatus
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		fmt.Println(string(debug.Stack()), "\n", consumerGroupID, "getConsumerGroup  NewClient : ", err)
		log.Error(ctx, consumerGroupID, "getConsumerGroup NewClient : ", err)
		os.Exit(1)
	}
	consumer, err := sarama.NewConsumerGroupFromClient(consumerGroupID, client)
	return &client, &consumer, err
}
