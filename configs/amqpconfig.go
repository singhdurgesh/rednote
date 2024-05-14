package configs

type AMQPConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	Exchange       string
	ExchangeType   string
	Queue          string
	BindingKey     string
	RoutingKey     string
	ConsumerTag    string
	WorkerPoolSize int
}
