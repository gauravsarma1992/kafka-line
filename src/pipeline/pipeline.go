package pipeline

type (
	LocationT       string
	TransformationT string
	UuidT           string

	Pipeline struct {
		Rules []*Rule `json:"rules"`
	}

	Rule struct {
		Source         *Location       `json:"source"`
		Destination    *Location       `json:"destination"`
		Transformation *Transformation `json:"transformation"`
		Filters        *Filters        `json:"filters"`
	}

	Transformation struct {
		Uuid               UuidT           `json:"uuid"`
		TransformationType TransformationT `json:"transformation_type"`
	}

	Filters struct {
		FilterType string `json:"filter_type"`
	}
)
