package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"jetsend_opens/shared/helpers"
	"jetsend_opens/shared/log"

	"github.com/Shopify/sarama"
)

var (
	eventProducer sarama.AsyncProducer
)

//InitializeKafa ...
func InitializeKafa(ctx context.Context) {
	// Setup configuration
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5

	// The level of acknowledgement reliability needed from the broker.
	config.Producer.RequiredAcks = sarama.WaitForAll

	var err error
	eventProducer, err = sarama.NewAsyncProducer(ServiceConfig.Kafka.Brokers, config)
	if err != nil {
		fmt.Println("InitializeKafa eventProducer err : ", err)
		log.Error(ctx, "InitializeKafa eventProducer err: ", err)
		return
	}
	// Track errors
	go func() {
		for {
			select {
			case err := <-eventProducer.Errors():
				msgData, _ := err.Msg.Value.Encode()
				log.Error(ctx, "Topic: ", err.Msg.Topic, " Data: ", string(msgData), "  kafka error: ", err.Error())
			}
		}
	}()

}

//writeToKafkaTopic function write MSG to Kafka Topic
func writeToKafkaTopic(ctx context.Context, open *helpers.Opens) error {

	event, err := helpers.GetEventMataData(helpers.Opened, *open)
	if err != nil {
		log.Error(ctx, "writeToKafkaTopic GetEventMataData err : ", err)
		return err
	}

	dataByte, err := json.Marshal(event)
	if err != nil {
		log.Error(ctx, event.AccountID, " writeToKafkaTopic json Marshal err : ", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: ServiceConfig.Kafka.EventTopic,
		//Key:       sarama.StringEncoder(strTime),
		Value:     sarama.StringEncoder(string(dataByte)),
		Timestamp: open.RecordedAT,
	}

	eventProducer.Input() <- msg

	return nil
}
