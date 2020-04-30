package lib

//ServiceConfig  Configuration Structs
var ServiceConfig Configuration

//Configuration ...
type Configuration struct {
	Kafka
	Service
}

type Kafka struct {
	OpensTopic string   `yaml:"opens_topic"`
	EventTopic string   `yaml:"event_topic"`
	Brokers    []string `yaml:"brokers"`
}

//KafkaConfiguration ...
type KafkaConfiguration struct {
	Env map[string]Kafka `yaml:"Environment"`
}

//Service ...
type Service struct {
	Port       string `yaml:"port"`
	ProjectID  string `yaml:"project_id"`
	LogName    string `yaml:"log_name"`
	PixelImage string `yaml:"pixel_image"`
}

//ServiceConfiguration ...
type ServiceConfiguration struct {
	Env map[string]Service `yaml:"Environment"`
}
