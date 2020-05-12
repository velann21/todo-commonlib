package async_jobs

import (
	"fmt"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	BOOTSTRAP_SERVER = "bootstrap.servers"
	IDEMPOTENCE = "enable.idempotence"
)
type Kafka struct {
	Broker string
	Topic string
	Producer *kafka.Producer
	Term chan bool
	Done chan bool
	
	ConfigMap *kafka.ConfigMap
	ConsumerObj *kafka.Consumer
}

func New(broker string, topic string) *Kafka{
	return &Kafka{Broker:broker, Topic:topic}
}

func (kf *Kafka) KafkaConfig() *Kafka{
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		BOOTSTRAP_SERVER: kf.Broker,
		IDEMPOTENCE: true})

	kf.Producer = p
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return kf
}

func (kf *Kafka) WithAsync() *Kafka {
	go func() {
		doTerm := false
		for !doTerm {
			select {
			case e := <-kf.Producer.Events():
				switch ev := e.(type) {
				case *kafka.Message:
					// Message delivery report
					m := ev
					if m.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
					} else {
						fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
							*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
					}

				case kafka.Error:
					e := ev
					if e.IsFatal() {
						fmt.Printf("FATAL ERROR: %v: terminating\n", e)
					} else {
						fmt.Printf("Error: %v\n", e)
					}

				default:
					fmt.Printf("Ignored event: %s\n", ev)
				}

			case <-kf.Term:
				doTerm = true
			}
		}

		close(kf.Done)
	}()
	return kf
}

func (kf *Kafka) Push(key []byte, value []byte) error{
	err := kf.Producer.Produce(&kafka.Message{
		Key:key,
		TopicPartition: kafka.TopicPartition{Topic: &kf.Topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		return err
	}
	kf.Producer.Flush(15 * 1000)
	return nil
}


func (kf *Kafka) BuildProducer(kafkaProducer Kafka,key []byte, value []byte) error{
	err := kf.Push(key, value)
	if err != nil{
		return err
	}
	return nil
}





