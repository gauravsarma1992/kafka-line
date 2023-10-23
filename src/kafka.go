package pipeline

type (
	KafkaServer struct {
		Config *KafkaConfig
	}
	KafkaConfig struct {
		BootstrapServers []string `json:"bootstrap_servers"`
	}
)

func NewKafkaServer(config *KafkaConfig) (kSrv *KafkaServer, err error) {
	kSrv = &KafkaServer{
		Config: config,
	}
	return
}
