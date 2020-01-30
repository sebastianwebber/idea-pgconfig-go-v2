package config

// OutputType defines the parameter output type
type OutputType int

const (
	// Text formats the output as a regular string
	Text OutputType = iota

	// Bytes formats the output using human bytes syntax
	Bytes

	// Number formats the output using human bytes syntax
	Number
)
