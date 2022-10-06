package sample

import (
	"encoding/gob"
	"fmt"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
)

// Fields return the list of extractor fields exported by this plugin.
// This method is mandatory the field extraction capability.
// If the Fields method is defined, the framework expects an Extract method
// to be specified too.
func (p *Plugin) Fields() []sdk.FieldEntry {
	return []sdk.FieldEntry{
		{
			// we currently support uint64 and string
			Type:    "uint64",
			Name:    "sample.field",
			Display: "Sample field",
			Desc:    "An example of field that this plugin can extract",
			// fields can extract a single value, or a list of values
			IsList: false,
			// fields can have an argument
			Arg: sdk.FieldEntryArg{
				IsRequired: false,
			},
		},
	}
}

// This method is mandatory the field extraction capability.
// If the Extract method is defined, the framework expects an Fields method
// to be specified too.
func (p *Plugin) Extract(req sdk.ExtractRequest, evt sdk.EventReader) error {
	var value uint64
	encoder := gob.NewDecoder(evt.Reader())
	if err := encoder.Decode(&value); err != nil {
		return err
	}

	switch req.Field() {
	case "sample.field":
		// once extracted, we can set the value in the extraction request
		req.SetValue(value)
	default:
		return fmt.Errorf("unsupported field: %s", req.Field())
	}

	return nil
}
