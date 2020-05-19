package async_jobs

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type ConsumerConfig struct {
	Server         []string
	AddressVersion string
	GroupID        string
	PartitionEOF   bool
	Offset         string
	Timeout        int
}

func NewConsumerConfig() *ConsumerConfig {
	return &ConsumerConfig{}
}

func (consumerConfig *ConsumerConfig) SetServer(setServer func()) *ConsumerConfig {
	setServer()
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

func (kf *Kafka) Consumer(serverString string, groupID string, chn chan interface{}) *Kafka {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	c, err := kafka.NewConsumer(kf.ConfigMap)
	if err != nil {
		log.Fatal("Error Occured")
	}
	kf.ConsumerObj = c
	return kf
}

func (kf *Kafka) AddSubscribers(func()) *Kafka {
	err := kf.ConsumerObj.SubscribeTopics([]string{"crawler-data"}, nil)
	if err != nil {
		log.Fatal("")
	}
	return kf
}

func (kf *Kafka) BuildConsumer(sigchan chan os.Signal, cdc chan []byte) {
	run := true
	rv := reflect.ValueOf(cdc)
	if rk := rv.Kind(); rk != reflect.Chan {
		panic("expecting type: 'chan ...'  instead got: " + rk.String())
	}
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
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	_ = kf.ConsumerObj.Close()
}