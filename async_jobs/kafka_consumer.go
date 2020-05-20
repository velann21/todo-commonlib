package async_jobs

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ConsumerConfig struct {
	Server         string
	AddressVersion string
	GroupID        string
	PartitionEOF   bool
	Offset         string
	Timeout        int
}

func NewConsumerConfig() *ConsumerConfig {
	return &ConsumerConfig{}
}

func (consumerConfig *ConsumerConfig) SetServer(setServer string) *ConsumerConfig {
	consumerConfig.Server = setServer
	return consumerConfig
}

func (consumerConfig *ConsumerConfig) SetAddress(addressVersion string) *ConsumerConfig {
	consumerConfig.AddressVersion = addressVersion
	return consumerConfig
}

func (consumerConfig *ConsumerConfig) SetGroupID(groupID string) *ConsumerConfig {
	consumerConfig.GroupID = groupID
	return consumerConfig
}

func (consumerConfig *ConsumerConfig) SetOffset(timeout int) *ConsumerConfig {
	consumerConfig.Timeout = timeout
	return consumerConfig
}

func (consumerConfig *ConsumerConfig) SetOffsetPtr(offsetPtr string) *ConsumerConfig {
	consumerConfig.Offset = offsetPtr
	return consumerConfig
}

func (kf *Kafka) KafkaConsumerConfig(cc *ConsumerConfig) *Kafka {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":     cc.Server,
		"broker.address.family": cc.AddressVersion,
		"group.id":              cc.GroupID,
		"session.timeout.ms":    cc.Timeout,
		"enable.partition.eof":  cc.PartitionEOF,
		"auto.offset.reset":     cc.Offset}
	kf.ConfigMap = configMap
	return kf
}

func (kf *Kafka) Consumer(sigchan chan os.Signal, cdc chan []byte) *Kafka {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"broker.address.family": "v4",
		"group.id":              "crawler_data_grp",
		"session.timeout.ms":    6000,
		"enable.partition.eof": true,
		"auto.offset.reset":     "earliest"})
	if err != nil {
		fmt.Print(err)
		log.Fatal("Error Occured")
	}
	err = c.SubscribeTopics([]string{"index-data"}, nil)
	if err != nil {
		log.Fatal("")
	}

	kf.ConsumerController = true
	run := kf.ConsumerController

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				cdc <- e.Value
			case *kafka.PartitionEOF:
				fmt.Println("Partion is EOF")
			case *kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Print("Error occured in kafka.Error")
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			}
		}
	}

	kf.ConsumerObj = c
	return kf
}

func (kf *Kafka) AddSubscribers(topics []string) *Kafka {
	err := kf.ConsumerObj.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatal("")
	}
	return kf
}

func (kf *Kafka) BuildConsumer(sigchan chan os.Signal, cdc chan []byte) {
	kf.ConsumerController = true
	run := kf.ConsumerController

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := kf.ConsumerObj.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				cdc <- e.Value
			case *kafka.PartitionEOF:
				fmt.Println("Partion is EOF")
			case *kafka.Error:
			// Errors should generally be considered
			// informational, the client will try to
			// automatically recover.
			// But in this example we choose to terminate
			// the application if all brokers are down.
			fmt.Print("Error occured in kafka.Error")
			fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		}
	}
	}

	fmt.Printf("Closing consumer\n")
	_ = kf.ConsumerObj.Close()
}

func (kf *Kafka) SendTerminationSignal(sigchan chan os.Signal){
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
}