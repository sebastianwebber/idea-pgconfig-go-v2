package config

const (
	// connMin contains the minimum value for the max_connections parameter
	connMin = 5
)

func maxConnections(i *Input) Parameter {

	out := Parameter{
		input: i,
		Name:  "max_connections",
		Type:  Number,
	}

	out.value = connMin

	return out
}
