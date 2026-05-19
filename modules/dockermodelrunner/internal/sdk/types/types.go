package types

import (
	"time"
)

// ModelResponse is the response for a model
type ModelResponse struct {
	// ID is the ID of the model
	ID string `json:"id"`

	// Tags are the tags of the model
	Tags []string `json:"tags"`

	// Config is the config of the model
	Config map[string]string `json:"config"`

	// Created is the creation time of the model
	Created time.Time `json:"created"`
}

// modelResponseAlias is an alias for ModelResponse to avoid recursion in JSON marshaling/unmarshaling.
// This is necessary because we want ModelResponse to contain a time.Time field which is not directly
// compatible with JSON serialization/deserialization.
type modelResponseAlias ModelResponse

// modelResponseJSON is a struct used for JSON marshaling/unmarshaling of ModelResponse.
// It includes a Unix timestamp for the Created field to ensure compatibility with JSON.
type modelResponseJSON struct {
	modelResponseAlias
	CreatedAt int64 `json:"created"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (mr *ModelResponse) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (mr ModelResponse) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
