package trackers

import "fmt"

type Factory func(cfg any) (Tracker, error)

var registry = map[string]Factory{}

func Register(trackerType string, factory Factory) {
	registry[trackerType] = factory
}

func New(trackerType string, cfg any) (Tracker, error) {
	factory, ok := registry[trackerType]
	if !ok {
		return nil, fmt.Errorf("unknown tracker type: %s", trackerType)
	}
	return factory(cfg)
}
