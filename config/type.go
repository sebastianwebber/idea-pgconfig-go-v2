package config

// OutputType defines the parameter output type
type OutputType int

const (
	// Bytes formats the output using human bytes syntax
	Bytes OutputType = iota

	// Text formats the output as a regular string
	Text
)
