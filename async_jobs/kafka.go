package async_jobs

type AsyncJobs interface {
	KafkaConfig() *Kafka
	WithAsync() *Kafka
	Push(key []byte, value []byte) error
	BuildProducer(kafkaProducer Kafka,key []byte, value []byte) error

	KafkaConsumerConfig() *Kafka
	KafkaConsumerSubscriber() *Kafka
	BuildConsumer() error

}
