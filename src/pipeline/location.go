package pipeline

type (
	Location struct {
		Uuid         UuidT     `json:"uuid"`
		LocationType LocationT `json:"location_type"`
	}
	KafkaTopicLocation struct {
		Name string `json:"name"`
	}
)

func NewKafkaLocation(topicName string) (kLoc *KafkaTopicLocation, err error) {
	return
}
