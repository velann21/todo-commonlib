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

func (kf *Kafka) Consumer() *Kafka {
	c, err := kafka.NewConsumer(kf.ConfigMap)
	if err != nil {
		fmt.Print(err)
		log.Fatal("Error Occured")
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

func (kf *Kafka) SendTerminationSignal(sigchan chan os.Signal){
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
}