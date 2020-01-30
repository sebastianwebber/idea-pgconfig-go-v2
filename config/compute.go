package config

// Compute does the magic
func (i *Input) Compute() *[]Category {
	var output []Category

	if i.PostgresVersion == 0 {
		i.PostgresVersion = PGVersion
	}

	memory := Category{
		Name: "Resource Usage / Memory",
	}
	memory.Parameters = append(memory.Parameters,
		sharedBuffers(i),
		maxConnections(i))

	output = append(output, memory)
	return &output
}

type Category struct {
	Name       string `json:"name"`
	Parameters []Parameter
}
