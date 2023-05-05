package config

type KafkaConfig struct {
	Host string
	Port int
}

func DefaultKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Host: "localhost",
		Port: 9092,
	}
}
