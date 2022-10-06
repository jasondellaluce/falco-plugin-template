package sample

import (
	"context"
	"encoding/gob"
	"strconv"
	"time"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
)

// Open opens the plugin source and starts a new capture session (e.g. stream
// of events), creating a new plugin instance. The state of each instance can
// be initialized here. This method is mandatory for the event sourcing capability.
func (m *Plugin) Open(params string) (source.Instance, error) {
	// we parse the user-passed parameters. For simplicity, we ignore the errors
	// and just start counting by 0 if the param is badly formatted
	start, err := strconv.Atoi(params)
	if err != nil || start < 0 {
		start = 0
	}

	// in this example, we simply write an integer counter in each event
	counter := uint64(start)
	pull := func(ctx context.Context, evt sdk.EventWriter) error {
		counter++
		if err := gob.NewEncoder(evt.Writer()).Encode(counter); err != nil {
			return err
		}
		evt.SetTimestamp(uint64(time.Now().UnixNano()))
		return nil
	}

	// event source capture sessions (defined by the source.Instance type) can
	// be implemented with either a functional paradigm (source.NewPullInstance),
	// a channel-based paradigm (source.NewPushInstance), or by defining
	// a custom type implementing source.Instance (advanced use cases only)
	return source.NewPullInstance(pull)
}
